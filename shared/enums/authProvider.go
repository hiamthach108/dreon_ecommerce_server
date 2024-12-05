package enums

type AuthProvider string

const (
	EmailPasswordAuthProvider AuthProvider = "EMAIL_PASSWORD"
	GoogleAuthProvider        AuthProvider = "GOOGLE"
	AppleAuthProvider         AuthProvider = "APPLE"
	FacebookAuthProvider      AuthProvider = "FACEBOOK"
)
