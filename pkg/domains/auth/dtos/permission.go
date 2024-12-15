package dtos

type PermissionDto struct {
	Id       string `json:"id,omitempty"`
	Resource string `json:"resource,omitempty"`
	Path     string `json:"path,omitempty"`
	Action   string `json:"action,omitempty"`
}
