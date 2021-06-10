package main

import (
	"context"
	"fmt"
	"go-base-learning/global"
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
	consulRegistry := consul.NewRegistry(registry.Addrs(global.GetConsulInfo()))

	svc := micro.NewService(
		micro.Name("student.client"),
		micro.Registry(consulRegistry))

	var bench [100]BenchTime
	index := autoAddIndex()
	now := time.Now()
	svc.Init()
	bench[index()] = BenchTime{"init server", time.Since(now)}

	studentService := message.NewStudentService("lb.student.endpoint", svc.Client())
	bench[index()] = BenchTime{"new client", time.Since(now)}

	res, err := studentService.GetStudent(context.TODO(), &message.StudentRequest{Name: "davie"})
	bench[index()] = BenchTime{"call server", time.Since(now)}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Name)
	fmt.Println(res.Classes)
	fmt.Println(res.Grade)
	//time.Sleep(50 * time.Second)

	for index, item := range bench {
		if item.value != 0 {
			fmt.Println(index, item.name, item.value.Microseconds())
		}
	}
}
