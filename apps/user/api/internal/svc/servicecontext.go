package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-demo/apps/user/api/internal/config"
	"go-zero-demo/apps/user/rpc/user"
)

// 注册服务上下文rpc 的依赖

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
