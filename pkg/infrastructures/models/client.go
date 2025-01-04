package models

import (
	"dreon_ecommerce_server/shared/enums"
	"time"

	"gorm.io/gorm"
)

type Client struct {
	Id          string              `json:"id" gorm:"primary_key;type:uuid"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Key         string              `json:"key,omitempty"`
	Secret      string              `json:"secret,omitempty"`
	Status      enums.GeneralStatus `json:"status,omitempty"`
	CreatedBy   string              `json:"createdBy,omitempty" gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
