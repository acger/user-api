package svc

import (
	"github.com/acger/user-api/internal/config"
	"github.com/acger/user-svc/userclient"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserSvc userclient.User

}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserSvc: userclient.NewUser(zrpc.MustNewClient(c.UserSvc)),
	}
}
