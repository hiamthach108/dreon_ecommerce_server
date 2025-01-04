package usecases

import (
	"context"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	"dreon_ecommerce_server/pkg/domains/auth/services"
	sharedI "dreon_ecommerce_server/shared/interfaces"
	"errors"
)

type authUsecase struct {
	IAuthUsecase
	logger  sharedI.ILogger
	userSvc services.IUserSvc
	authSvc services.IAuthSvc
}

type IAuthUsecase interface {
	Login(ctx context.Context, req *dtos.LoginReq) (result *dtos.LoginResp, err error)
	Register(ctx context.Context, req *dtos.RegisterReq) (result *dtos.RegisterResp, err error)
}

func NewAuthUsecase(logger sharedI.ILogger) *authUsecase {
	userSvc := services.NewUserSvc()
	authSvc := services.NewAuthSvc()

	return &authUsecase{
		logger:  logger,
		userSvc: userSvc,
		authSvc: authSvc,
	}
}

func (u *authUsecase) Login(ctx context.Context, req *dtos.LoginReq) (result *dtos.LoginResp, err error) {
	action := "authUsecase.Login"

	u.logger.Infof("%s start: email %s", action, req.Email)

	result = &dtos.LoginResp{}

	return result, nil
}

func (u *authUsecase) Register(ctx context.Context, req *dtos.RegisterReq) (result *dtos.RegisterResp, err error) {
	action := "authUsecase.Register"

	u.logger.Infof("%s start: email %s", action, req.Email)

	existedUser, err := u.userSvc.IsExistUserByEmail(ctx, req.Email)
	if err == nil && existedUser {
		return nil, errors.New("user existed")
	}

	user, err := u.authSvc.RegisterByPassword(ctx, req.Email, req.Password, req.FirstName, req.LastName)
	if err != nil || user == nil {
		u.logger.Errorf("[%s] error: %v", action, err)
		return nil, err
	}
	result = &dtos.RegisterResp{
		UserId: user.Id,
	}

	return result, nil
}
