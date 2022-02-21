package svc

import (
	"github.com/acger/user-api/internal/config"
	"github.com/acger/user-svc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserSvc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserSvc: user.NewUser(zrpc.MustNewClient(c.UserSvc)),
	}
}
