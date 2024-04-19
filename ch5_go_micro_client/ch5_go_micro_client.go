package main

import (
	"context"
	"fmt"
	"go-base-learning/global"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

	"github.com/micro/go-micro"

	message "go-base-learning/endpoint"
)

type BenchTime struct {
	name  string
	value time.Duration
}

func autoAddIndex() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// 创建 Consul 注册中心
	consulRegistry := consul.NewRegistry(
		registry.Addrs(global.GetConsulInfo()),
	)

	// 创建微服务
	svc := micro.NewService(
		micro.Name("student.client"),
		micro.Registry(consulRegistry),
	)

	// 初始化微服务
	svc.Init()

	// 创建学生服务客户端
	studentService := message.NewStudentService("dev.student.endpoint", svc.Client())

	// 创建一个信号接收器，捕获中断信号
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	// 创建一个 context ，用于跟踪信号
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 在新的 goroutine 中等待中断信号
	go func() {
		<-stopChan // 当接收到信号时，停止等待
		fmt.Println("Received stop signal, shutting down...")

		// 执行清理工作，关闭微服务
		svc.Server().Stop()
		cancel() // 取消 context，告诉其他 goroutine 退出
	}()

	// 调用学生服务的 GetStudent 方法
	res, err := studentService.GetStudent(ctx, &message.StudentRequest{Name: "davie"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印学生信息
	fmt.Println(res.Name)
	fmt.Println(res.Classes)
	fmt.Println(res.Grade)
}
