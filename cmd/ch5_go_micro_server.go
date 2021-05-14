package main

import (
	"context"
	"errors"
	"fmt"
	message "go-base-learning/endpoint"
	"log"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

	"github.com/micro/go-micro"
)

//学生服务管理实现
type StudentManager struct {
}

//获取学生信息的服务接口实现
func (sm *StudentManager) GetStudent(ctx context.Context, request *message.StudentRequest, response *message.Student) error {

	//tom
	studentMap := map[string]message.Student{
		"davie":  message.Student{Name: "davie", Classes: "软件工程专业", Grade: 80},
		"steven": message.Student{Name: "steven", Classes: "计算机科学与技术", Grade: 90},
		"tony":   message.Student{Name: "tony", Classes: "计算机网络工程", Grade: 85},
		"jack":   message.Student{Name: "jack", Classes: "工商管理", Grade: 96},
	}

	if request.Name == "" {
		return errors.New(" 请求参数错误,请重新请求。")
	}

	//获取对应的student
	student := studentMap[request.Name]

	if student.Name != "" {
		fmt.Println(student.Name, student.Classes, student.Grade)
		*response = student
		return nil
	}

	return errors.New(" 未查询当相关学生信息 ")
}

const (
	svcName = `lb.student.endpoint`
)

func main() {

	consulRegistry := consul.NewRegistry(registry.Addrs("localhost:8500"))

	//创建一个新的服务对象实例
	service := micro.NewService(
		micro.Name(svcName),
		micro.Version("1.0.0"),
		micro.Registry(consulRegistry))

	service.Init()

	message.RegisterStudentServiceHandler(service.Server(), new(StudentManager))

	err := service.Run()
	if err != nil {
		log.Fatal(err)
	}
}
