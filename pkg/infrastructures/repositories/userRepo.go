package repositories

import (
	"context"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	repoI "dreon_ecommerce_server/pkg/domains/auth/interfaces"
	"dreon_ecommerce_server/pkg/infrastructures/models"
	"dreon_ecommerce_server/shared/constants"
	"dreon_ecommerce_server/shared/enums"
	"dreon_ecommerce_server/shared/interfaces"

	"github.com/golobby/container/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	db     *gorm.DB
	logger interfaces.ILogger
	repoI.IUserRepo
}

func NewUserRepo() *userRepo {
	var (
		logger interfaces.ILogger
		db     *gorm.DB
	)
	container.Resolve(&logger)
	container.Resolve(&db)

	return &userRepo{
		db:     db,
		logger: logger,
	}
}

func (r *userRepo) FindAllUser(ctx context.Context, page, pageSize *int32, status *enums.UserStatus, search *string) (result *[]models.User, total int64, err error) {
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

	return
}

func (r *userRepo) FindUserByEmail(ctx context.Context, email string) (result *models.User, err error) {
	err = r.db.Where("email = ?", email).First(&result).Error
	return
}

func (r *userRepo) FindUserById(ctx context.Context, userId string) (result *models.User, err error) {
	var data models.User
	err = r.db.Where("id = ?", userId).First(&data).Error
	if err != nil {
		return
	}

	result = &data
	return
}

func (r *userRepo) ExistUserByEmail(ctx context.Context, email string) (result bool, err error) {
	var data models.User
	err = r.db.Where("email = ?", email).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return
	}

	return true, nil
}

func (r *userRepo) GetUserAuth(ctx context.Context, userId string) (result *[]models.UserAuth, err error) {
	var data []models.UserAuth
	err = r.db.Where("user_id = ?", userId).Find(&data).Error
	if err != nil {
		return
	}

	result = &data
	return
}

func (r *userRepo) Create(ctx context.Context, user *models.User) (result *models.User, err error) {
	user.Id = uuid.New().String()

	err = r.db.Create(user).Error
	if err != nil {
		return
	}

	result = user
	return
}

func (r *userRepo) UpdateLastLogin(ctx context.Context, userId string) (err error) {
	err = r.db.Model(&models.User{}).Where("id = ?", userId).Update("last_login", "NOW()").Error
	return
}

func (r *userRepo) UpdateStatus(ctx context.Context, userId string, status enums.UserStatus) (err error) {
	err = r.db.Model(&models.User{}).Where("id = ?", userId).Update("status", status).Error
	return
}

func (r *userRepo) UpdateUser(ctx context.Context, user *models.User) (result *models.User, err error) {
	err = r.db.Save(user).Error
	if err != nil {
		return
	}

	result = user
	return
}

func (r *userRepo) UpsertUserAuth(ctx context.Context, userAuth *dtos.UserAuthDto) (result *models.UserAuth, err error) {
	err = r.db.Where("user_id = ? AND client_id = ? AND role_id = ?", userAuth.UserId, userAuth.ClientId, userAuth.RoleId).FirstOrCreate(userAuth).Error
	if err != nil {
		return
	}

	result = &models.UserAuth{
		UserId:   userAuth.UserId,
		ClientId: userAuth.ClientId,
		RoleId:   userAuth.RoleId,
	}
	return
}

func (r *userRepo) DeleteUserAuth(ctx context.Context, userId, clientId, roleId string) (err error) {
	err = r.db.Where("user_id = ? AND client_id = ? AND role_id = ?", userId, clientId, roleId).Delete(&models.UserAuth{}).Error
	return
}
