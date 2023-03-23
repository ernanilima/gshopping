package repository_test

import (
	"reflect"
	"testing"

	"github.com/ernanilima/gshopping/app/config"
	"github.com/ernanilima/gshopping/app/repository"
	brand_repository "github.com/ernanilima/gshopping/app/repository/brand"
	"github.com/stretchr/testify/assert"
)

// Deve retornar o total de metodos existentes
func TestRepository_Should_Return_The_Total_Methods(t *testing.T) {
	totalMethods := reflect.TypeOf((*repository.Repository)(nil)).Elem().NumMethod()
	assert.Equal(t, 3, totalMethods)
}

// Deve retornar os metodos para a interface brand repository
func TestNewRepository_Should_Return_Methods_For_BrandRepository(t *testing.T) {
	controller := repository.NewRepository(&config.Config{})

	brandController, exist := controller.(brand_repository.BrandRepository)
	typeOf := reflect.TypeOf(brandController)
	assert.True(t, exist)
	assert.NotNil(t, typeOf)

	nameOfMethods := []string{
		"FindAll",
		"FindByDescription",
		"FindById",
	}

	for index, name := range nameOfMethods {
		assert.Equal(t, name, typeOf.Method(index).Name)
	}
}
