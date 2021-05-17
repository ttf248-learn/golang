package main

import (
	"context"
	"go-base-learning/global"
	"log"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"

	pb "go-base-learning/endpoint"

	"github.com/micro/go-micro"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello" + req.Name
	return nil
}

func main() {
	consulRegister := consul.NewRegistry(registry.Addrs(global.GetConsulInfo()))

	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Registry(consulRegister))

	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
