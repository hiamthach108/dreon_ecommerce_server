package dtos

import "dreon_ecommerce_server/shared/enums"

type ClientDto struct {
	Id          string              `json:"id"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Key         string              `json:"key,omitempty"`
	Secret      string              `json:"secret,omitempty"`
	Status      enums.GeneralStatus `json:"status,omitempty"`
	CreatedBy   string              `json:"createdBy,omitempty"`
}

type PublicClientDto struct {
	Id          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
