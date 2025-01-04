package dtos

type UserDto struct {
	Id        string `json:"id"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	BirthDate int64  `json:"birthDate,omitempty"`
	LastLogin int64  `json:"lastLogin,omitempty"`
}

type UserAuthDto struct {
	UserId   string `json:"userId"`
	ClientId string `json:"clientId"`
	RoleId   string `json:"roleIds"`
}
