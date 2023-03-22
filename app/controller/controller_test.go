package controller_test

import (
	"reflect"
	"testing"

	"github.com/ernanilima/gshopping/app/config"
	"github.com/ernanilima/gshopping/app/controller"
	brand_controller "github.com/ernanilima/gshopping/app/controller/brand"
	"github.com/ernanilima/gshopping/app/repository"
	"github.com/stretchr/testify/assert"
)

// Deve retornar o total de metodos existentes
func TestController_Should_Return_The_Total_Methods(t *testing.T) {
	totalMethods := reflect.TypeOf((*controller.Controller)(nil)).Elem().NumMethod()
	assert.Equal(t, 3, totalMethods)
}

// Deve retornar os metodos para a interface brand controller
func TestNewController_Should_Return_Methods_For_BrandController(t *testing.T) {
	controller := controller.NewController(repository.NewRepository(&config.Config{}))

	brandController, exist := controller.(brand_controller.BrandController)
	typeOf := reflect.TypeOf(brandController)
	assert.True(t, exist)
	assert.NotNil(t, typeOf)

	nameOfMethods := []string{
		"FindAllBrands",
		"FindAllBrandsByDescription",
		"FindBrandById",
	}

	for index, name := range nameOfMethods {
		assert.Equal(t, name, typeOf.Method(index).Name)
	}
}
