package models

import (
	"dreon_ecommerce_server/shared/enums"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id          string              `json:"id" gorm:"primary_key;type:uuid"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Status      enums.GeneralStatus `json:"status,omitempty"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RolePermission struct {
	RoleId       string `json:"roleId"`
	PermissionId string `json:"permissionId"`
}
