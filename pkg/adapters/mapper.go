package adapters

import (
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	"dreon_ecommerce_server/pkg/infrastructures/models"
	"log"

	"github.com/devfeel/mapper"
	"github.com/golobby/container/v3"
)

func IoCMapper() {
	container.Singleton(func() mapper.IMapper {
		var (
			mapperProvider = mapper.NewMapper()
			err            error
		)
		mapperProvider.SetEnableFieldIgnoreTag(true)
		defer func() {
			if err != nil {
				panic(err)
			}
		}()

		//register dtos
		err = mapperProvider.Register(&dtos.UserDto{})
		err = mapperProvider.Register(&dtos.UserAuthDto{})
		err = mapperProvider.Register(&dtos.ClientDto{})
		err = mapperProvider.Register(&dtos.PublicClientDto{})
		err = mapperProvider.Register(&dtos.RoleDto{})
		err = mapperProvider.Register(&dtos.PermissionDto{})
		err = mapperProvider.Register(&dtos.RolePermissionDto{})

		//register models
		err = mapperProvider.Register(&models.User{})
		err = mapperProvider.Register(&models.UserAuth{})
		err = mapperProvider.Register(&models.Client{})
		err = mapperProvider.Register(&models.Role{})
		err = mapperProvider.Register(&models.Permission{})
		err = mapperProvider.Register(&models.RolePermission{})

		if err != nil {
			log.Fatal("Error register mapper: ", err)
		}

		return mapperProvider
	})
}
