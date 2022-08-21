package main

import (
	"context"
	"flag"
	"fmt"

	"zero-demo-study/user-rpc/internal/config"
	"zero-demo-study/user-rpc/internal/server"
	"zero-demo-study/user-rpc/internal/svc"
	"zero-demo-study/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserCenterServer(grpcServer, server.NewUserCenterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	//添加拦截器
	s.AddUnaryInterceptors(TestServerInterceptor)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func TestServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("-------------before interceptor-------------\n")
	resp, err = handler(ctx, req)
	fmt.Printf("%v\n", resp)
	fmt.Printf("----------after interceptor----------------\n")
	return
}
