package repositories

import (
	"context"
	repoI "dreon_ecommerce_server/pkg/domains/auth/interfaces"
	"dreon_ecommerce_server/pkg/infrastructures/models"
	"dreon_ecommerce_server/shared/constants"
	"dreon_ecommerce_server/shared/enums"
	"dreon_ecommerce_server/shared/helpers"
	sharedI "dreon_ecommerce_server/shared/interfaces"

	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

type permissionRepo struct {
	db     *gorm.DB
	logger sharedI.ILogger
	repoI.IPermissionRepo
}

func NewPermissionRepo() *permissionRepo {
	var (
		logger sharedI.ILogger
		db     *gorm.DB
	)
	container.Resolve(&logger)
	container.Resolve(&db)

	return &permissionRepo{
		db:     db,
		logger: logger,
	}
}

func (r *permissionRepo) GetAllPermissions(ctx context.Context, page, pageSize *int32, search *string) (result []*models.Permission, total int64, err error) {
	query := r.db.Model(&models.Permission{})
	if search != nil {
		query = query.Where("name LIKE ?", "%"+*search+"%")
	}

	if page == nil {
		page = new(int32)
		*page = 0
	}
	if pageSize == nil {
		pageSize = new(int32)
		*pageSize = constants.DEFAULT_PAGE_SIZE
	}

	err = query.Count(&total).Error
	if err != nil {
		return
	}

	err = query.Offset(int(*page * *pageSize)).Limit(int(*pageSize)).Find(&result).Error
	if err != nil {
		return
	}

	return
}

func (r *permissionRepo) GetPermissionById(ctx context.Context, permissionId string) (result *models.Permission, err error) {
	var data models.Permission
	err = r.db.Where("id = ?", permissionId).First(&data).Error
	if err != nil {
		return
	}

	return &data, nil
}

func (r *permissionRepo) CreatePermission(ctx context.Context, permission *[]models.Permission) (result int32, err error) {
	for i := range *permission {
		(*permission)[i].Id = helpers.GenerateUUID()
	}

	err = r.db.Create(permission).Error
	if err != nil {
		return
	}

	return int32(len(*permission)), nil
}

func (r *permissionRepo) UpdatePermission(ctx context.Context, permission *models.Permission) (result *models.Permission, err error) {
	err = r.db.Save(permission).Error
	if err != nil {
		return
	}

	return permission, nil
}

func (r *permissionRepo) UpdateStatus(ctx context.Context, permissionId string, status enums.GeneralStatus) (err error) {
	err = r.db.Model(&models.Permission{}).Where("id = ?", permissionId).Update("status", status).Error
	if err != nil {
		return
	}

	return nil
}
