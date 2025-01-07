package repositories

import (
	"context"
	repoI "dreon_ecommerce_server/pkg/domains/auth/interfaces"
	"dreon_ecommerce_server/pkg/infrastructures/models"
	"dreon_ecommerce_server/shared/helpers"
	sharedI "dreon_ecommerce_server/shared/interfaces"

	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

type clientRepo struct {
	db     *gorm.DB
	logger sharedI.ILogger
	repoI.IClientRepo
}

func NewClientRepo() *clientRepo {
	var (
		logger sharedI.ILogger
		db     *gorm.DB
	)
	container.Resolve(&logger)
	container.Resolve(&db)

	return &clientRepo{
		db:     db,
		logger: logger,
	}
}

func (r *clientRepo) GetClientById(ctx context.Context, clientId string) (result *models.Client, err error) {
	err = r.db.Where("id = ?", clientId).First(&result).Error
	if err != nil {
		return
	}

	return
}

func (r *clientRepo) GetAllClients(ctx context.Context, page, pageSize *int32, search *string) (result []*models.Client, total int64, err error) {
	query := r.db.Model(&models.Client{})
	if search != nil {
		query = query.Where("name LIKE ?", "%"+*search+"%")
	}

	if page == nil {
		page = new(int32)
		*page = 0
	}
	if pageSize == nil {
		pageSize = new(int32)
		*pageSize = 10
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

func (r *clientRepo) CreateClient(ctx context.Context, client *models.Client) (result *models.Client, err error) {
	client.Id = helpers.GenerateUUID()
	err = r.db.Create(client).Error
	if err != nil {
		return
	}

	return client, nil
}

func (r *clientRepo) UpdateClient(ctx context.Context, client *models.Client) (result *models.Client, err error) {
	err = r.db.Save(client).Error
	if err != nil {
		return
	}

	return client, nil
}

func (r *clientRepo) UpdateStatus(ctx context.Context, clientId string, status bool) (err error) {
	err = r.db.Model(&models.Client{}).Where("id = ?", clientId).Update("status", status).Error
	if err != nil {
		return
	}

	return
}
