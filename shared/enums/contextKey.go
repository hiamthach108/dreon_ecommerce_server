package enums

type ContextKey string

const (
	AuthPayloadContextKey ContextKey = "Auth"
	UserIDContextKey      ContextKey = "UserID"
	EmailContextKey       ContextKey = "Email"
)
