package models

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	Id        string `json:"id" gorm:"primary_key;type:uuid"`
	Resource  string `json:"resource,omitempty"`
	Path      string `json:"path,omitempty"`
	Action    string `json:"action,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
