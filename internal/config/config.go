package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	UserSvc zrpc.RpcClientConf

	Qiniu struct{
		Bucket string
		AK string
		SK string
	}
}
