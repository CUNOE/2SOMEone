package svc

import (
	"github.com/leaper-one/2SOMEone/api/user-api/internal/config"
	"github.com/leaper-one/2SOMEone/rpc/user-rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	User   user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   user.NewUser(zrpc.MustNewClient(c.User)),
	}
}
