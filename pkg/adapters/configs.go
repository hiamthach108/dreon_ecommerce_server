package adapters

import (
	"dreon_ecommerce_server/configs"

	"github.com/golobby/container/v3"
)

func IocConfigs() {
	container.Singleton(func() *configs.AppConfig {
		appConfig, err := configs.NewAppConfig(".env")
		if err != nil {
			panic(err)
		}
		return appConfig
	})
}
