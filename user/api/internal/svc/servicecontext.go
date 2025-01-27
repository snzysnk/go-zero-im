package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-im/api/internal/config"
	"go-zero-im/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
