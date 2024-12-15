package models

import "dreon_ecommerce_server/shared/enums"

type Client struct {
	BaseModel   `json:",inline"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Key         string              `json:"key,omitempty"`
	Secret      string              `json:"secret,omitempty"`
	Status      enums.GeneralStatus `json:"status,omitempty"`
	CreatedBy   string              `json:"createdBy,omitempty" gorm:"not null"`
}
