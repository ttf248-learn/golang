package main

import (
	"flag"
	"fmt"

	"go-base-learning/ch14_go_zero_rpc/ch14_go_zero_rpc"
	"go-base-learning/ch14_go_zero_rpc/internal/config"
	"go-base-learning/ch14_go_zero_rpc/internal/server"
	"go-base-learning/ch14_go_zero_rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/ch14gozerorpc.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		ch14_go_zero_rpc.RegisterCh14GoZeroRpcServer(grpcServer, server.NewCh14GoZeroRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
