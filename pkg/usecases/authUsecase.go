package usecases

import (
	"context"
	"dreon_ecommerce_server/configs"
	"dreon_ecommerce_server/libs/jwt"
	"dreon_ecommerce_server/libs/oauth"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	"dreon_ecommerce_server/pkg/domains/auth/services"
	"dreon_ecommerce_server/shared/constants"
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
	oauth   oauth.IAppOAuth
	mapper  mapper.IMapper
	cache   sharedI.ICache
	userSvc services.IUserSvc
	authSvc services.IAuthSvc
}

type IAuthUsecase interface {
	Login(ctx context.Context, req *dtos.LoginReq) (result *dtos.LoginResp, err error)
	Register(ctx context.Context, req *dtos.RegisterReq) (result *dtos.RegisterResp, err error)
	GetUserProfile(ctx context.Context, userId string) (result *dtos.UserDto, err error)
	RefreshToken(ctx context.Context, req *dtos.RefreshTokenReq) (result *dtos.RefreshTokenResp, err error)
	GoogleOAuthCallBack(ctx context.Context, code, state string) (result *dtos.LoginResp, err error)
}

func NewAuthUsecase(appConfigs *configs.AppConfig, logger sharedI.ILogger) *authUsecase {
	userSvc := services.NewUserSvc()
	authSvc := services.NewAuthSvc()

	var mapperProvider mapper.IMapper
	container.Resolve(&mapperProvider)
	var cache sharedI.ICache
	container.Resolve(&cache)

	oauthLib := oauth.NewAppOAuth(appConfigs, logger)

	return &authUsecase{
		logger:  logger,
		configs: appConfigs,
		mapper:  mapperProvider,
		userSvc: userSvc,
		authSvc: authSvc,
		cache:   cache,
		oauth:   oauthLib,
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
		url, state := u.oauth.GetGoogleOAuthUrl()
		result.Google = &dtos.OAuthGoogleResp{
			Url:   url,
			State: state,
		}
		return result, nil

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

	accessToken, accessTokenExpAt, refreshToken, refreshTokenExpAt, err := u.generateToken(user, false)
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

	accessToken, accessTokenExpAt, refreshToken, refreshTokenExpAt, err := u.generateToken(user, false)
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

func (u *authUsecase) generateToken(user *dtos.UserDto, isRefresh bool) (accessToken string, accessTokenExpAt int64, refreshToken string, refreshTokenExpAt int64, err error) {
	action := "authUsecase.generateToken"

	jwt := jwt.NewEchoJWT(u.configs.Auth.JWT.SecretKey, u.configs.Auth.JWT.Issuer, u.mapper, u.configs.Auth.IgnoreMethods)

	accessTokenExpAt = time.Now().Add(time.Duration(u.configs.Auth.JWT.ExpiredTime) * time.Minute).UTC().Unix()
	accessToken, err = jwt.GenToken(user.Id, user.Email, enums.EmailPasswordAuthenType, accessTokenExpAt)
	if err != nil {
		u.logger.Errorf("[%s] error: %v", action, err)

		return
	}

	if !isRefresh {
		refreshDuration := time.Duration(u.configs.Auth.JWT.RefreshExpired) * time.Minute
		refreshTokenExpAt = time.Now().Add(refreshDuration).UTC().Unix()
		refreshToken, err = helpers.GenerateRandomString(constants.REFRESH_TOKEN_LEN)
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
	}

	return
}

func (u *authUsecase) GetUserProfile(ctx context.Context, userId string) (result *dtos.UserDto, err error) {
	action := "authUsecase.GetUserProfile"

	u.logger.Infof("%s start: userId %s", action, userId)

	user, err := u.userSvc.FindUserById(ctx, userId)
	if err != nil {
		u.logger.Errorf("[%s] error: %v", action, err)
		return
	}

	return user, nil
}

func (u *authUsecase) RefreshToken(ctx context.Context, req *dtos.RefreshTokenReq) (result *dtos.RefreshTokenResp, err error) {
	action := "authUsecase.RefreshToken"

	u.logger.Infof("%s start: refreshToken %s", action, req.RefreshToken, req.UserId)

	if req.RefreshToken == "" {
		return nil, errors.New("refresh token is required")
	}

	var user *dtos.UserDto

	tokenLen := len(req.RefreshToken)

	if tokenLen == constants.REFRESH_TOKEN_OAUTH_LEN {
		state := fmt.Sprintf("users:state:%s", req.RefreshToken)
		userData, err := u.cache.Get(state)
		if err != nil {
			u.logger.Errorf("[%s] error: %v", action, err)
			return nil, err
		}

		if userData == nil {
			return nil, errors.New("state not found")
		}

		user, err = u.userSvc.FindUserById(ctx, (userData).(string))
		if err != nil {
			u.logger.Errorf("[%s] error: %v", action, err)
			return nil, err
		}
	} else {
		userId, err := u.cache.Get(fmt.Sprintf("users:%s:rfk:%s", req.UserId, req.RefreshToken))
		if err != nil {
			u.logger.Errorf("[%s] error: %v", action, err)
			return nil, err
		}

		if userId == nil {
			return nil, errors.New("refresh token not found")
		}

		user, err = u.userSvc.FindUserById(ctx, (userId).(string))
		if err != nil {
			u.logger.Errorf("[%s] error: %v", action, err)
			return nil, err
		}
	}

	accessToken, accessTokenExpAt, _, _, err := u.generateToken(user, true)
	if err != nil {
		u.logger.Errorf("[%s] error: %v", action, err)
		return nil, err
	}

	result = &dtos.RefreshTokenResp{
		AccessToken:    accessToken,
		AccessTokenExp: accessTokenExpAt,
	}

	return
}

func (u *authUsecase) GoogleOAuthCallBack(ctx context.Context, code, state string) (result *dtos.LoginResp, err error) {
	action := "authUsecase.GoogleOAuthCallBack"

	u.logger.Infof("%s start: code %s", action, code)

	googleUser, err := u.oauth.GetGoogleUserInfo(ctx, code)
	if err != nil {
		u.logger.Errorf("[%s] error: %v", action, err)
		return
	}

	result = &dtos.LoginResp{}

	user, err := u.userSvc.FindUserByEmail(ctx, googleUser.Email)
	if err != nil || user == nil {
		newUser := &dtos.UserDto{
			Email:     googleUser.Email,
			FirstName: googleUser.GivenName,
			LastName:  googleUser.FamilyName,
			LastLogin: time.Now().UTC().Unix(),
		}

		user, err = u.userSvc.CreateOAuthUser(ctx, newUser, enums.GoogleAuthenType)
		if err != nil || user == nil {
			u.logger.Errorf("[%s] error: %v", action, err)
			return nil, err
		}
	}
	result.UserId = user.Id

	// set state to cache like refresh token
	key := fmt.Sprintf("users:%s:rfk:%s", user.Id, state)
	refreshDuration := time.Duration(u.configs.Auth.JWT.RefreshExpired) * time.Minute
	err = u.cache.Set(key, user.Id, &refreshDuration)
	if err != nil {
		u.logger.Errorf("[%s] error: %v", action, err)
		return
	}

	// set user data state
	stateKey := fmt.Sprintf("users:state:%s", state)
	err = u.cache.Set(stateKey, user.Id, &refreshDuration)
	if err != nil {
		u.logger.Errorf("[%s] error: %v", action, err)
		return
	}

	return
}
