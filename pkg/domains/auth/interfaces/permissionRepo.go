package interfaces

import (
	"context"
	"dreon_ecommerce_server/pkg/infrastructures/models"
	"dreon_ecommerce_server/shared/enums"
)

type IPermissionRepo interface {
	GetAllPermissions(ctx context.Context, page, pageSize *int32, search *string) (result []*models.Permission, total int64, err error)
	GetPermissionById(ctx context.Context, permissionId string) (result *models.Permission, err error)

	CreatePermission(ctx context.Context, permission *[]models.Permission) (result int32, err error)
	UpdatePermission(ctx context.Context, permission *models.Permission) (result *models.Permission, err error)
	UpdateStatus(ctx context.Context, permissionId string, status enums.GeneralStatus) (err error)
}
