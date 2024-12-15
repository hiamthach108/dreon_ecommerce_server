package models

import (
	"time"

	"gorm.io/gorm"
)

// base gorm model
type BaseModel struct {
	Id        string `gorm:"primary_key;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
