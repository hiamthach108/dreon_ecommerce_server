package adapters

import (
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

		if err != nil {
			log.Fatal("Error register mapper: ", err)
		}

		return mapperProvider
	})
}
