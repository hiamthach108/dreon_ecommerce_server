package services

import (
	"context"
	"dreon_ecommerce_server/libs/crypto"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	"dreon_ecommerce_server/pkg/domains/auth/interfaces"
	"dreon_ecommerce_server/pkg/infrastrutures/models"
	"dreon_ecommerce_server/pkg/infrastrutures/repositories"
	"dreon_ecommerce_server/shared/enums"
	sharedI "dreon_ecommerce_server/shared/interfaces"
	"errors"

	"github.com/devfeel/mapper"
	"github.com/golobby/container/v3"
)

type authSvc struct {
	logger   sharedI.ILogger
	userRepo interfaces.IUserRepo
	mapper   mapper.IMapper
	crypto   crypto.IPasswordEncoder
}

type IAuthSvc interface {
	RegisterByPassword(ctx context.Context, email, password, firstName, lastName string) (user *dtos.UserDto, err error)
	LoginByPassword(ctx context.Context, email, password string) (user *dtos.UserDto, err error)
}

func NewAuthSvc() *authSvc {
	var (
		logger sharedI.ILogger
		mapper mapper.IMapper
		crypto crypto.IPasswordEncoder
	)
	container.Resolve(&logger)
	container.Resolve(&mapper)
	container.Resolve(&crypto)

	userRepo := repositories.NewUserRepo()

	return &authSvc{
		logger:   logger,
		userRepo: userRepo,
		mapper:   mapper,
		crypto:   crypto,
	}
}

func (s *authSvc) RegisterByPassword(ctx context.Context, email, password, firstName, lastName string) (user *dtos.UserDto, err error) {
	action := "authSvc.RegisterByPassword"

	hashedPassword, err := s.crypto.Hash(password)
	if err != nil {
		s.logger.Errorf("[%s] error %v", action, err)
		return
	}

	newUser := models.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Password:  hashedPassword,
		AuthType:  string(enums.EmailPasswordAuthenType),
		Status:    enums.USER_STATUS_ACTIVE,
	}

	new, err := s.userRepo.Create(ctx, &newUser)
	if err != nil {
		s.logger.Errorf("[%s] error %v", action)
		return
	}

	user = &dtos.UserDto{
		Id: new.Id,
	}

	err = s.mapper.AutoMapper(new, user)

	return
}

func (s *authSvc) LoginByPassword(ctx context.Context, email, password string) (user *dtos.UserDto, err error) {
	action := "authSvc.LoginByPassword"

	u, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil || u == nil {
		s.logger.Errorf("[%s] error %v", action, err)
		return nil, errors.New("user not found")
	}

	err = s.crypto.Compare(password, u.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	user = &dtos.UserDto{}

	s.mapper.Mapper(u, user)

	return user, nil
}
