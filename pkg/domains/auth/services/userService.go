package services

import (
	"context"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	"dreon_ecommerce_server/pkg/domains/auth/interfaces"
	"dreon_ecommerce_server/shared/enums"
	sharedI "dreon_ecommerce_server/shared/interfaces"
)

type userSvc struct {
	logger   sharedI.ILogger
	userRepo interfaces.IUserRepo
}

func NewUserSvc(logger sharedI.ILogger, userRepo interfaces.IUserRepo) *userSvc {
	return &userSvc{
		logger:   logger,
		userRepo: userRepo,
	}
}

func (s *userSvc) FindAllUser(ctx context.Context, page, pageSize *int32, status *enums.UserStatus, search *string) (result *[]dtos.UserDto, total int64, err error) {
	action := "userSvc.FindAllUser"

	result, total, err = s.userRepo.FindAllUser(ctx, page, pageSize, status, search)
	if err != nil {
		s.logger.Errorf("%s[findAll] error %v", action, err)
		return
	}

	return
}
