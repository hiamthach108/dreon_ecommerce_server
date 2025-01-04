package test

import (
	"dreon_ecommerce_server/pkg/adapters"
	"dreon_ecommerce_server/pkg/domains/auth/dtos"
	"dreon_ecommerce_server/pkg/infrastrutures/models"
	"testing"

	"github.com/devfeel/mapper"
	"github.com/golobby/container/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapper_ModelToDto(t *testing.T) {
	adapters.IoCMapper()

	var mapperProvider mapper.IMapper
	err := container.Resolve(&mapperProvider)
	require.NoError(t, err, "error resolve mapperProvider")
	assert.NotNil(t, mapperProvider, "mapperProvider is nil")

	userModel := &models.User{
		Id:        "testId",
		Email:     "test@gmail.com",
		FirstName: "test",
		LastName:  "test",
		BirthDate: 0,
		LastLogin: 0,
	}

	userDto := &dtos.UserDto{
		Id: userModel.Id,
	}

	err = mapperProvider.Mapper(userModel, userDto)
	require.NoError(t, err, "error mapping model to dto")
}

func TestMapper_DtoToModel(t *testing.T) {
	adapters.IoCMapper()

	var mapperProvider mapper.IMapper
	err := container.Resolve(&mapperProvider)
	require.NoError(t, err, "error resolve mapperProvider")
	assert.NotNil(t, mapperProvider, "mapperProvider is nil")

	userDto := &dtos.UserDto{
		Id:        "abc",
		Email:     "test@gmail.com",
		FirstName: "test",
		LastName:  "test",
		BirthDate: 0,
		LastLogin: 0,
	}

	userModel := &models.User{}

	err = mapperProvider.Mapper(userDto, userModel)
	require.NoError(t, err, "error mapping dto to model")

	assert.Equal(t, userDto.Email, userModel.Email, "Email should be equal")
	assert.Equal(t, userDto.FirstName, userModel.FirstName, "FirstName should be equal")
	assert.Equal(t, userDto.LastName, userModel.LastName, "LastName should be equal")
	assert.Equal(t, userDto.BirthDate, userModel.BirthDate, "BirthDate should be equal")
	assert.Equal(t, userDto.LastLogin, userModel.LastLogin, "LastLogin should be equal")
}

func TestMapper_NilValues(t *testing.T) {
	adapters.IoCMapper()

	var mapperProvider mapper.IMapper
	err := container.Resolve(&mapperProvider)
	require.NoError(t, err, "error resolve mapperProvider")
	assert.NotNil(t, mapperProvider, "mapperProvider is nil")

	var userModel *models.User
	userDto := &dtos.UserDto{}

	err = mapperProvider.Mapper(userModel, userDto)
	assert.Error(t, err, "mapping nil model should produce an error")

	userModel = &models.User{}
	var userDtoNil *dtos.UserDto

	err = mapperProvider.Mapper(userModel, userDtoNil)
	assert.Error(t, err, "mapping to nil dto should produce an error")
}
