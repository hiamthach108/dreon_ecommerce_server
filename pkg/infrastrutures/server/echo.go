package server

import (
	"dreon_ecommerce_server/configs"
	"dreon_ecommerce_server/libs/jwt"
	"fmt"
	"net/http"
	"time"

	"github.com/devfeel/mapper"
	"github.com/golobby/container/v3"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
)

func StartEchoServer() {
	var (
		configs        = &configs.AppConfig{}
		e              = echo.New()
		mapperProvider mapper.IMapper
		err            error
	)

	err = container.Resolve(&configs)
	if err != nil {
		panic(err)
	}

	jwtEcho := jwt.NewEchoJWT(configs.Auth.JWT.SecretKey, configs.Auth.JWT.Issuer, mapperProvider, configs.Auth.IgnoreMethods)

	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Content-Type", "application/json;charset=UTF-8")
			return next(c)
		}
	})
	e.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: jwtEcho.NewClaimFunc,
		SigningKey:    jwtEcho.GetSignedKey(),
		Skipper:       jwtEcho.Skipper,
		ErrorHandler: func(c echo.Context, err error) error {
			return nil
		},
		ContinueOnIgnoredError: true,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlMaxAge,
			echo.HeaderAcceptEncoding,
			echo.HeaderAccessControlAllowCredentials,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderCacheControl,
			echo.HeaderContentLength,
			echo.HeaderUpgrade,
		},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	}))

	var h2s = &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          30 * time.Second,
	}
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.Logger.Error(e.StartH2CServer(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port), h2s))

}
