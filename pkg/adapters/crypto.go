package adapters

import (
	"dreon_ecommerce_server/libs/crypto"

	"github.com/golobby/container/v3"
)

func IoCCrypto() {
	container.Singleton(func() crypto.IPasswordEncoder {
		return crypto.NewPasswordEncoder()
	})
}
