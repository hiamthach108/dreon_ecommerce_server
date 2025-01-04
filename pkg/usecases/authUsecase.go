package usecases

import (
	"context"
	"dreon_ecommerce_server/configs"
	"dreon_ecommerce_server/libs/jwt"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	"dreon_ecommerce_server/pkg/domains/auth/services"
	"dreon_ecommerce_server/shared/enums"
	"dreon_ecommerce_server/shared/helpers"
	sharedI "dreon_ecommerce_server/shared/interfaces"
	"errors"
	"fmt"
	"time"

	"github.com/devfeel/mapper"
	"github.com/golobby/container/v3"
)

type authUsecase struct {
	IAuthUsecase
	logger  sharedI.ILogger
	configs *configs.AppConfig
	mapper  mapper.IMapper
	cache   sharedI.ICache
	userSvc services.IUserSvc
	authSvc services.IAuthSvc
}

type IAuthUsecase interface {
	Login(ctx context.Context, req *dtos.LoginReq) (result *dtos.LoginResp, err error)
	Register(ctx context.Context, req *dtos.RegisterReq) (result *dtos.RegisterResp, err error)
}

func NewAuthUsecase(appConfigs *configs.AppConfig, logger sharedI.ILogger) *authUsecase {
	userSvc := services.NewUserSvc()
	authSvc := services.NewAuthSvc()

	var mapperProvider mapper.IMapper
	container.Resolve(&mapperProvider)
	var cache sharedI.ICache
	container.Resolve(&cache)

	return &authUsecase{
		logger:  logger,
		configs: appConfigs,
		mapper:  mapperProvider,
		userSvc: userSvc,
		authSvc: authSvc,
		cache:   cache,
	}
}

func (u *authUsecase) Login(ctx context.Context, req *dtos.LoginReq) (result *dtos.LoginResp, err error) {
	action := "authUsecase.Login"

	u.logger.Infof("%s start: email %s", action, req.Email)
	result = &dtos.LoginResp{}
	var user *dtos.UserDto

	switch req.AuthenType {
	case enums.EmailPasswordAuthenType:
		user, err = u.authSvc.LoginByPassword(ctx, req.Email, req.Password)
		if err != nil || user == nil {
			u.logger.Errorf("[%s] error: %v", action, err)
			return nil, err
		}
	case enums.GoogleAuthenType:
		// TODO: implement google login
		return nil, errors.New("google login not implemented")
	case enums.FacebookAuthenType:
		// TODO: implement facebook login
		return nil, errors.New("facebook login not implemented")
	case enums.AppleAuthenType:
		// TODO: implement apple login
		return nil, errors.New("apple login not implemented")

	default:
		return nil, errors.New("authen type not supported")
	}

	result.UserId = user.Id

	accessToken, accessTokenExpAt, refreshToken, refreshTokenExpAt, err := u.generateToken(user)
	if err != nil {
		u.logger.Errorf("[%s] error: %v", action, err)
		return nil, err
	}

	result.AccessToken = accessToken
	result.AccessTokenExp = accessTokenExpAt
	result.RefreshToken = refreshToken
	result.RefreshTokenExp = refreshTokenExpAt

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

	accessToken, accessTokenExpAt, refreshToken, refreshTokenExpAt, err := u.generateToken(user)
	if err != nil {
		u.logger.Errorf("[%s] error: %v", action, err)
		return result, nil
	}

	result.AccessToken = accessToken
	result.AccessTokenExp = accessTokenExpAt
	result.RefreshToken = refreshToken
	result.RefreshTokenExp = refreshTokenExpAt

	return result, nil
}

func (u *authUsecase) generateToken(user *dtos.UserDto) (accessToken string, accessTokenExpAt int64, refreshToken string, refreshTokenExpAt int64, err error) {
	action := "authUsecase.generateToken"

	jwt := jwt.NewEchoJWT(u.configs.Auth.JWT.SecretKey, u.configs.Auth.JWT.Issuer, u.mapper, u.configs.Auth.IgnoreMethods)

	accessTokenExpAt = time.Now().Add(time.Duration(u.configs.Auth.JWT.ExpiredTime) * time.Minute).UTC().Unix()
	accessToken, err = jwt.GenToken(user.Id, user.Email, enums.EmailPasswordAuthenType, accessTokenExpAt)
	if err != nil {
		u.logger.Errorf("[%s] error: %v", action, err)

		return
	}

	refreshDuration := time.Duration(u.configs.Auth.JWT.RefreshExpired) * time.Minute
	refreshTokenExpAt = time.Now().Add(refreshDuration).UTC().Unix()
	refreshToken, err = helpers.GenerateRandomString(24)
	if err != nil {
		u.logger.Errorf("[%s] error: %v", action, err)
		return
	}

	cacheKey := fmt.Sprintf("users:%s:rfk:%s", user.Id, refreshToken)
	err = u.cache.Set(cacheKey, user.Id, &refreshDuration)
	if err != nil {
		u.logger.Errorf("[%s] error: %v", action, err)
		return
	}

	return
}
