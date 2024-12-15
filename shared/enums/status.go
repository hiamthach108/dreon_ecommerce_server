package enums

type UserStatus string

const (
	USER_STATUS_ACTIVE   UserStatus = "ACTIVE"
	USER_STATUS_INACTIVE UserStatus = "INACTIVE" // user registered but not yet verify in email
	USER_STATUS_BLOCKED  UserStatus = "BLOCKED"  // user blocked by admin
	USER_STATUS_DELETED  UserStatus = "DELETED"  // user deleted by admin
)

type GeneralStatus string

const (
	STATUS_ACTIVE   GeneralStatus = "ACTIVE"
	STATUS_INACTIVE GeneralStatus = "INACTIVE"
	STATUS_DELETED  GeneralStatus = "DELETED"
	STATUS_ARCHIVED GeneralStatus = "ARCHIVED"
)
