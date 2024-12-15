package dtos

import "dreon_ecommerce_server/shared/enums"

type UserDto struct {
	Id        string           `json:"id"`
	Email     string           `json:"username,omitempty"`
	FirstName string           `json:"firstName,omitempty"`
	LastName  string           `json:"lastName,omitempty"`
	Password  string           `json:"password,omitempty"`
	BirthDate int64            `json:"birthDate,omitempty"`
	Status    enums.UserStatus `json:"status,omitempty"`
	AuthType  enums.AuthenType `json:"authType,omitempty"`
	LastLogin int64            `json:"lastLogin,omitempty"`
}

type UserAuthDto struct {
	UserId   string `json:"userId"`
	ClientId string `json:"clientId"`
	RoleId   string `json:"roleIds"`
}
