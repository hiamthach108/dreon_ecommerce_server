package enums

type Role string

const (
	UserRole      Role = "USER"
	AdminFullRole Role = "ADMIN_FULL"
	SemiAdmin     Role = "SEMI_ADMIN"
)

type Permission string

const (
	ReadPermission   Permission = "READ"
	CreatePermission Permission = "CREATE"
	UpdatePermission Permission = "UPDATE"
	DeletePermission Permission = "DELETE"
)

type Resource string

const (
	UserDataResource Resource = "USER"
	WordDataResource Resource = "WORD"
)
