package enums

type Role string

const (
	UserRole     Role = "USER"
	SupperAdmin  Role = "SUPERADMIN"
	Admin        Role = "ADMIN"
	ClientAdmin  Role = "CLIENT_ADMIN"
	ClientUser   Role = "CLIENT_USER"
	ClientGuest  Role = "CLIENT_GUEST"
	ClientEditor Role = "CLIENT_EDITOR"
)

type Permission string

const (
	ReadPermission       Permission = "READ"
	CreatePermission     Permission = "CREATE"
	UpdatePermission     Permission = "UPDATE"
	ManagementPermission Permission = "MANAGEMENT"
)

type Resource string

const (
	// Auth
	UserResource       Resource = "USER"
	ClientResource     Resource = "CLIENT"
	ClientUserResource Resource = "CLIENT_USER"
	RoleResource       Resource = "ROLE"
	PermissionResource Resource = "PERMISSION"

	// Product
	ProductResource  Resource = "PRODUCT"
	BrandResource    Resource = "BRAND"
	CategoryResource Resource = "CATEGORY"
	OrderResource    Resource = "ORDER"
	PaymentResource  Resource = "PAYMENT"
	ReportResource   Resource = "REPORT"

	// Other
	SettingResource  Resource = "SETTING"
	AnalyticResource Resource = "ANALYTIC"
)
