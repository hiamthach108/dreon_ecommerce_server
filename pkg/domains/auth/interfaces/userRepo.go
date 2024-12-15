package interfaces

import (
	"context"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	"dreon_ecommerce_server/shared/enums"
)

type IUserRepo interface {
	FindAllUser(ctx context.Context, page, pageSize *int32, status *enums.UserStatus, search *string) (result *[]dtos.UserDto, total int64, err error)
	FindUserByEmail(ctx context.Context, email string) (result *dtos.UserDto, err error)
	FindUserById(ctx context.Context, userId string) (result *dtos.UserDto, err error)
	ExistUserByEmail(ctx context.Context, email string) (result bool, err error)
	GetUserAuth(ctx context.Context, userId string) (result *[]dtos.UserAuthDto, err error)

	Create(ctx context.Context, user *dtos.UserDto) (result *dtos.UserDto, err error)
	UpdateLastLogin(ctx context.Context, userId string) (err error)
	UpdateStatus(ctx context.Context, userId string, status enums.UserStatus) (err error)
	UpdateUser(ctx context.Context, user *dtos.UserDto) (result *dtos.UserDto, err error)

	// auth
	UpsertUserAuth(ctx context.Context, userAuth *dtos.UserAuthDto) (result *dtos.UserAuthDto, err error)
	DeleteUserAuth(ctx context.Context, userId, clientId, roleId string) (err error)
}
