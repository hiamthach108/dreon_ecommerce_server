package services

import (
	"context"
	"dreon_ecommerce_server/libs/crypto"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	"dreon_ecommerce_server/pkg/domains/auth/interfaces"
	"dreon_ecommerce_server/pkg/infrastrutures/repositories"
	"dreon_ecommerce_server/shared/enums"
	sharedI "dreon_ecommerce_server/shared/interfaces"

	"github.com/devfeel/mapper"
	"github.com/golobby/container/v3"
)

type userSvc struct {
	logger   sharedI.ILogger
	userRepo interfaces.IUserRepo
	mapper   mapper.IMapper
	crypto   crypto.IPasswordEncoder
}

type IUserSvc interface {
	FindAllUser(ctx context.Context, page, pageSize *int32, status *enums.UserStatus, search *string) (result *[]dtos.UserDto, total int64, err error)
	IsExistUserByEmail(ctx context.Context, email string) (result bool, err error)
}

func NewUserSvc() *userSvc {
	var (
		logger sharedI.ILogger
		mapper mapper.IMapper
		crypto crypto.IPasswordEncoder
	)
	container.Resolve(&logger)
	container.Resolve(&mapper)
	container.Resolve(&crypto)

	userRepo := repositories.NewUserRepo()

	return &userSvc{
		logger:   logger,
		userRepo: userRepo,
		mapper:   mapper,
		crypto:   crypto,
	}
}

func (s *userSvc) FindAllUser(ctx context.Context, page, pageSize *int32, status *enums.UserStatus, search *string) (result *[]dtos.UserDto, total int64, err error) {
	action := "userSvc.FindAllUser"

	r, total, err := s.userRepo.FindAllUser(ctx, page, pageSize, status, search)
	if err != nil {
		s.logger.Errorf("%s error %v", action, err)
		return
	}

	err = s.mapper.Mapper(&r, &result)
	if err != nil {
		return
	}

	return
}

func (s *userSvc) IsExistUserByEmail(ctx context.Context, email string) (result bool, err error) {
	action := "userSvc.IsExistUserByEmail"

	result, err = s.userRepo.ExistUserByEmail(ctx, email)
	if err != nil {
		s.logger.Errorf("%s error %v", action, err)
		return
	}

	return
}
