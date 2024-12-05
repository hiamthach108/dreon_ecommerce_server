package adapters

import (
	"dreon_ecommerce_server/libs/logger"
	"dreon_ecommerce_server/shared/interfaces"

	"github.com/golobby/container/v3"
)

func IoCLogger() {
	container.SingletonLazy(func() interfaces.ILogger {
		log, err := logger.NewAppLogger()
		if err != nil {
			panic(err)
		}

		return log
	})
}
