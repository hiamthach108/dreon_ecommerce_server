package models

type Permission struct {
	BaseModel `json:",inline"`
	Resource  string `json:"resource,omitempty"`
	Path      string `json:"path,omitempty"`
	Action    string `json:"action,omitempty"`
}
