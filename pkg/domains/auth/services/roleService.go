package services

import (
	"dreon_ecommerce_server/pkg/domains/auth/interfaces"
	sharedI "dreon_ecommerce_server/shared/interfaces"

	"github.com/devfeel/mapper"
	"github.com/golobby/container/v3"
)

type roleSvc struct {
	logger         sharedI.ILogger
	mapper         mapper.IMapper
	roleRepo       interfaces.IRoleRepo
	permissionRepo interfaces.IPermissionRepo
}

type IRoleSvc interface {
}

func NewRoleSvc() *roleSvc {
	var (
		logger sharedI.ILogger
		mapper mapper.IMapper
	)
	container.Resolve(&logger)
	container.Resolve(&mapper)

	return &roleSvc{
		logger: logger,
		mapper: mapper,
	}
}
