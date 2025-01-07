package repositories

import (
	"context"
	repoI "dreon_ecommerce_server/pkg/domains/auth/interfaces"
	"dreon_ecommerce_server/pkg/infrastructures/models"
	"dreon_ecommerce_server/shared/enums"
	"dreon_ecommerce_server/shared/helpers"
	sharedI "dreon_ecommerce_server/shared/interfaces"

	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

type roleRepo struct {
	db     *gorm.DB
	logger sharedI.ILogger
	repoI.IRoleRepo
}

func NewRoleRepo() *roleRepo {
	var (
		logger sharedI.ILogger
		db     *gorm.DB
	)
	container.Resolve(&logger)
	container.Resolve(&db)

	return &roleRepo{
		db:     db,
		logger: logger,
	}
}

func (r *roleRepo) GetAllRoles(ctx context.Context, clientId *string) (result []*models.Role, err error) {
	query := r.db.Model(&models.Role{})
	if clientId != nil {
		query = query.Where("client_id = ?", *clientId)
	}

	var data []*models.Role
	err = query.Find(&data).Error
	if err != nil {
		return
	}

	return data, nil
}

func (r *roleRepo) GetRoleById(ctx context.Context, roleId string) (result *models.Role, err error) {
	var data models.Role
	err = r.db.Where("id = ?", roleId).First(&data).Error
	if err != nil {
		return
	}

	return &data, nil
}

func (r *roleRepo) GetRoleByName(ctx context.Context, roleName string) (result *models.Role, err error) {
	var data models.Role
	err = r.db.Where("name = ?", roleName).First(&data).Error
	if err != nil {
		return
	}

	return &data, nil
}

func (r *roleRepo) CreateRole(ctx context.Context, role *models.Role) (result *models.Role, err error) {
	role.Id = helpers.GenerateUUID()

	err = r.db.Create(role).Error
	if err != nil {
		return
	}

	return role, nil
}

func (r *roleRepo) UpdateRole(ctx context.Context, role *models.Role) (result *models.Role, err error) {
	err = r.db.Save(role).Error
	if err != nil {
		return
	}

	return role, nil
}

func (r *roleRepo) UpdateStatus(ctx context.Context, roleId string, status enums.GeneralStatus) (err error) {
	err = r.db.Model(&models.Role{}).Where("id = ?", roleId).Update("status", status).Error
	if err != nil {
		return
	}

	return nil
}
