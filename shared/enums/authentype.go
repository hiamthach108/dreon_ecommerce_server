package enums

type AuthenType string

const (
	EmailPasswordAuthenType AuthenType = "EMAIL_PASSWORD"
	GoogleAuthenType        AuthenType = "GOOGLE"
	AppleAuthenType         AuthenType = "APPLE"
	FacebookAuthenType      AuthenType = "FACEBOOK"
)
