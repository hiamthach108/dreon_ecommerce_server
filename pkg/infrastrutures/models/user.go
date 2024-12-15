package models

import "dreon_ecommerce_server/shared/enums"

type User struct {
	BaseModel `json:",inline"`
	Email     string           `json:"username,omitempty" gorm:"unique;not null"`
	FirstName string           `json:"firstName,omitempty"`
	LastName  string           `json:"lastName,omitempty"`
	Password  string           `json:"password,omitempty"`
	BirthDate int64            `json:"birthDate,omitempty"`
	Status    enums.UserStatus `json:"status,omitempty"`
	AuthType  string           `json:"authType,omitempty"`
	LastLogin int64            `json:"lastLogin,omitempty"`
}

type UserAuth struct {
	UserId   string `json:"userId"`
	ClientId string `json:"clientId"`
	RoleId   string `json:"roleIds"`
}
