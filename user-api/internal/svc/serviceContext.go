package svc

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"zero-demo-study/user-api/internal/config"
	"zero-demo-study/user-api/internal/middleware"
	"zero-demo-study/user-rpc/usercenter"
)

type ServiceContext struct {
	Config         config.Config
	TestMiddleware rest.Middleware
	UserRpcClient  usercenter.UserCenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
		UserRpcClient:  usercenter.NewUserCenter(zrpc.MustNewClient(c.UserConf, zrpc.WithUnaryClientInterceptor(TestClientInterceptor))),
	}
}

func TestClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("client interceptor before")
	md := metadata.New(map[string]string{"username": "zhangsan"})
	ctx = metadata.NewOutgoingContext(ctx, md)
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}
	fmt.Println("client interceptor after")
	return nil
}
