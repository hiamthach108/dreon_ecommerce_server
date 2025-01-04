package dtos

import "dreon_ecommerce_server/shared/enums"

type LoginReq struct {
	Email      string           `json:"email"`
	Password   string           `json:"password"`
	AuthenType enums.AuthenType `json:"authenType"`
}

type LoginResp struct {
	UserId          string `json:"userId"`
	AccessToken     string `json:"accessToken"`
	RefreshToken    string `json:"refreshToken"`
	AccessTokenExp  int64  `json:"accessTokenExp"`
	RefreshTokenExp int64  `json:"refreshTokenExp"`
}

type RegisterReq struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type RegisterResp struct {
	UserId string `json:"userId"`
	LoginResp
}
