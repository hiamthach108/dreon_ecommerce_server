package adapters

import (
	"dreon_ecommerce_server/configs"
	"dreon_ecommerce_server/libs/cache"
	"dreon_ecommerce_server/shared/interfaces"
	"log"
	"time"

	"github.com/golobby/container/v3"
)

func IoCCache() {
	container.SingletonLazy(func() interfaces.ICache {
		var (
			logger    interfaces.ILogger
			appConfig *configs.AppConfig
			err       error
		)

		err = container.Resolve(&logger)
		if err != nil {
			log.Fatal(err)
		}
		err = container.Resolve(&appConfig)
		if err != nil {
			log.Fatal(err)
		}

		expiredTime := time.Duration(appConfig.Cache.DefaultExpireTimeSec) * time.Second
		cleanupInterval := time.Duration(appConfig.Cache.CleanupIntervalHour) * time.Hour
		return cache.NewAppCache(logger, &expiredTime, &cleanupInterval)
	})
}
