package models

import (
	"dreon_ecommerce_server/shared/enums"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        string           `json:"id" gorm:"primary_key;type:uuid"`
	Email     string           `json:"username,omitempty" gorm:"unique;not null"`
	FirstName string           `json:"firstName,omitempty"`
	LastName  string           `json:"lastName,omitempty"`
	Password  string           `json:"password,omitempty"`
	BirthDate int64            `json:"birthDate,omitempty"`
	LastLogin int64            `json:"lastLogin,omitempty"`
	Status    enums.UserStatus `json:"status,omitempty"`
	AuthType  string           `json:"authType,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserAuth struct {
	UserId   string `json:"userId"`
	ClientId string `json:"clientId"`
	RoleId   string `json:"roleIds"`
}
