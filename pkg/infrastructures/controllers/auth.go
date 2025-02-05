package controllers

import (
	"dreon_ecommerce_server/configs"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	"dreon_ecommerce_server/pkg/usecases"
	"net/http"

	"dreon_ecommerce_server/shared/constants"
	"dreon_ecommerce_server/shared/enums"
	"dreon_ecommerce_server/shared/helpers"
	sharedI "dreon_ecommerce_server/shared/interfaces"

	"github.com/labstack/echo/v4"
)

type authController struct {
	appConfigs  *configs.AppConfig
	logger      sharedI.ILogger
	authUsecase usecases.IAuthUsecase
}

func NewAuthController(appConfigs *configs.AppConfig, logger sharedI.ILogger) *authController {
	au := usecases.NewAuthUsecase(appConfigs, logger)

	return &authController{
		appConfigs:  appConfigs,
		logger:      logger,
		authUsecase: au,
	}
}

func (c *authController) Login(ctx echo.Context) (err error) {
	body := &dtos.LoginReq{}
	if err := ctx.Bind(body); err != nil {
		appErr := constants.NewBadRequest(err, "invalid request body")
		return appErr.ToEchoHTTPError()
	}

	result, err := c.authUsecase.Login(ctx.Request().Context(), body)
	if err != nil || result == nil {
		if appErr, ok := err.(*constants.AppError); ok {
			return appErr.ToEchoHTTPError()
		}
		appErr := constants.NewBadRequest(err, "login failed")

		return appErr.ToEchoHTTPError()
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *authController) Register(ctx echo.Context) (err error) {
	body := &dtos.RegisterReq{}
	if err := ctx.Bind(body); err != nil {
		appErr := constants.NewBadRequest(err, "invalid request body")
		return appErr.ToEchoHTTPError()
	}

	result, err := c.authUsecase.Register(ctx.Request().Context(), body)
	if err != nil || result == nil {
		if appErr, ok := err.(*constants.AppError); ok {
			return appErr.ToEchoHTTPError()
		}
		appErr := constants.NewBadRequest(err, "register failed")

		return appErr.ToEchoHTTPError()
	}

	return helpers.SuccessResponse(ctx, result)
}

func (c *authController) GetProfile(ctx echo.Context) (err error) {
	userID, ok := ctx.Request().Context().Value(enums.UserIDContextKey).(string)
	if !ok {
		appErr := constants.NewUnAuthorize(nil, "unauthorize")
		return appErr.ToEchoHTTPError()
	}

	result, err := c.authUsecase.GetUserProfile(ctx.Request().Context(), userID)
	if err != nil || result == nil {
		if appErr, ok := err.(*constants.AppError); ok {
			return appErr.ToEchoHTTPError()
		}
		appErr := constants.NewBadRequest(err, "get profile failed")

		return appErr.ToEchoHTTPError()
	}

	return helpers.SuccessResponse(ctx, result)
}

func (c *authController) RefreshToken(ctx echo.Context) (err error) {
	body := &dtos.RefreshTokenReq{}
	if err := ctx.Bind(body); err != nil {
		appErr := constants.NewBadRequest(err, "invalid request body")
		return appErr.ToEchoHTTPError()
	}

	result, err := c.authUsecase.RefreshToken(ctx.Request().Context(), body)
	if err != nil || result == nil {
		if appErr, ok := err.(*constants.AppError); ok {
			return appErr.ToEchoHTTPError()
		}
		appErr := constants.NewBadRequest(err, "refresh token failed")

		return appErr.ToEchoHTTPError()
	}

	return helpers.SuccessResponse(ctx, result)
}
