package svc

import (
	"zero-admin/internal/config"
	"zero-admin/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	Permission rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Permission: middleware.NewPermissionMiddleware().Handle,
	}
}
