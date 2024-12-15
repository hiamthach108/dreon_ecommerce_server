package adapters

import (
	"dreon_ecommerce_server/configs"
	"dreon_ecommerce_server/pkg/infrastrutures/models"
	"dreon_ecommerce_server/shared/interfaces"
	"log"

	"github.com/golobby/container/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func IoCDatabase() {
	container.Singleton(func() *gorm.DB {
		var (
			appConfig *configs.AppConfig
			logger    interfaces.ILogger
			err       error
		)

		err = container.Resolve(&appConfig)
		if err != nil {
			log.Fatal(err)
		}

		err = container.Resolve(&logger)
		if err != nil {
			log.Fatal(err)
		}

		db, err := gorm.Open(postgres.Open(appConfig.Postgres.Dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		logger.InfoF("Connected to database: %s", appConfig.Postgres.Dsn)

		db.AutoMigrate(
			&models.User{},
			&models.Client{},
			&models.Role{},
			&models.Permission{},
			&models.RolePermission{},
			&models.UserAuth{},
		)

		logger.Info("Migrated tables")

		return db
	})
}
