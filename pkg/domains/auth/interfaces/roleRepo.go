package interfaces

import (
	"context"
	"dreon_ecommerce_server/pkg/infrastructures/models"
	"dreon_ecommerce_server/shared/enums"
)

type IRoleRepo interface {
	GetAllRoles(ctx context.Context, clientId *string) (result []*models.Role, err error)
	GetRoleById(ctx context.Context, roleId string) (result *models.Role, err error)
	GetRoleByName(ctx context.Context, roleName string) (result *models.Role, err error)

	CreateRole(ctx context.Context, role *models.Role) (result *models.Role, err error)
	UpdateRole(ctx context.Context, role *models.Role) (result *models.Role, err error)
	UpdateStatus(ctx context.Context, roleId string, status enums.GeneralStatus) (err error)
}
