package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-im/api/internal/config"
	"go-zero-im/api/internal/middleware"
	"go-zero-im/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	userclient.User
	AuthInterceptor rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		////不需要额外穿参
		AuthInterceptor: middleware.NewAuthInterceptorMiddleware().Handle,
	}
}
