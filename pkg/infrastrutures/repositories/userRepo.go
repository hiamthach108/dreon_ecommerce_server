package repositories

import (
	"context"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	repoI "dreon_ecommerce_server/pkg/domains/auth/interfaces"
	"dreon_ecommerce_server/pkg/infrastrutures/models"
	"dreon_ecommerce_server/shared/constants"
	"dreon_ecommerce_server/shared/enums"
	"dreon_ecommerce_server/shared/interfaces"

	"github.com/devfeel/mapper"
	"gorm.io/gorm"
)

type userRepo struct {
	db     *gorm.DB
	logger interfaces.ILogger
	mapper mapper.IMapper
	repoI.IUserRepo
}

func NewUserRepo(db *gorm.DB, logger interfaces.ILogger, mapper mapper.IMapper) *userRepo {
	return &userRepo{
		db:     db,
		logger: logger,
		mapper: mapper,
	}
}

func (r *userRepo) FindAllUser(ctx context.Context, page, pageSize *int32, status *enums.UserStatus, search *string) (result *[]dtos.UserDto, total int64, err error) {
	if page == nil {
		page = new(int32)
		*page = 0
	}
	if pageSize == nil {
		pageSize = new(int32)
		*pageSize = constants.DEFAULT_PAGE_SIZE
	}

	var totalData int64

	query := r.db.Model(&models.User{})
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if search != nil {
		query = query.Where("email LIKE ?", "%"+*search+"%")
	}

	err = query.Count(&totalData).Error
	if err != nil {
		return
	}

	var data []models.User
	err = query.Offset(int(*page * *pageSize)).Limit(int(*pageSize)).Find(&data).Error
	if err != nil {
		return
	}

	err = r.mapper.Mapper(&data, &result)
	if err != nil {
		return
	}

	return
}
