package models

import "dreon_ecommerce_server/shared/enums"

type Role struct {
	BaseModel   `json:",inline"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Status      enums.GeneralStatus `json:"status,omitempty"`
}

type RolePermission struct {
	RoleId       string `json:"roleId"`
	PermissionId string `json:"permissionId"`
}
