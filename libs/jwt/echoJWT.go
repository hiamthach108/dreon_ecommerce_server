package jwt

import (
	"dreon_ecommerce_server/shared/constants"
	"dreon_ecommerce_server/shared/enums"
	"fmt"
	"time"

	"github.com/devfeel/mapper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type echoJWT struct {
	signedKey      string
	issuer         string
	skipper        map[string]bool
	mapperProvider mapper.IMapper
	IEchoJWT
}

type IEchoJWT interface {
	GenToken(id, email string, provider enums.AuthenType) (token string, err error)
	NewClaimFunc(c echo.Context) jwt.Claims
	Skipper(c echo.Context)
	GetSignedKey()
}

type JwtCustomClaim struct {
	Id         string           `json:"id,omitempty"`
	Email      string           `json:"email,omitempty"`
	AuthenType enums.AuthenType `json:"authenType,omitempty"`
	jwt.RegisteredClaims
}

func (j *JwtCustomClaim) Valid() error {
	if len(j.Id) <= 0 {
		return constants.NewUnAuthorize(fmt.Errorf("invalid payload"), "unauthorize")
	}
	if j.ExpiresAt.Before(time.Now().UTC()) {
		return constants.NewUnAuthorize(fmt.Errorf("token expired"), "unauthorize")
	}
	return nil
}

func NewEchoJWT(signedKey, issuer string, mapperProvider mapper.IMapper, skipper map[string]bool) *echoJWT {
	return &echoJWT{
		signedKey:      signedKey,
		issuer:         issuer,
		mapperProvider: mapperProvider,
		skipper:        skipper,
	}
}

func (e *echoJWT) GenToken(id, email string, provider enums.AuthenType) (token string, err error) {
	claim := &JwtCustomClaim{
		Id:         id,
		Email:      email,
		AuthenType: provider,
	}
	claim.ExpiresAt = jwt.NewNumericDate(time.Now().UTC().Add(time.Hour * 8640))
	claim.IssuedAt = jwt.NewNumericDate(time.Now().UTC())
	claim.Issuer = e.issuer
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err = jwtToken.SignedString([]byte(e.signedKey))
	return
}

func (e *echoJWT) NewClaimFunc(c echo.Context) jwt.Claims {
	return new(JwtCustomClaim)
}

func (e *echoJWT) Skipper(c echo.Context) bool {
	if e.skipper != nil && len(e.skipper) > 0 {
		return e.skipper[c.Path()]
	}
	return false
}

func (e *echoJWT) GetSignedKey() []byte {
	return []byte(e.signedKey)
}
