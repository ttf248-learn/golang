package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

	"github.com/micro/go-micro"

	message "go-base-learning/endpoint"
)

func main() {
	consulRegistry := consul.NewRegistry(registry.Addrs("localhost:8500"))

	svc := micro.NewService(
		micro.Name("student.client"),
		micro.Registry(consulRegistry))

	svc.Init()

	studentService := message.NewStudentService("lb.student.endpoint", svc.Client())

	res, err := studentService.GetStudent(context.TODO(), &message.StudentRequest{Name: "davie"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Name)
	fmt.Println(res.Classes)
	fmt.Println(res.Grade)
	time.Sleep(50 * time.Second)
}
