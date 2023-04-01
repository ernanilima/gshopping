package controller_test

import (
	"reflect"
	"testing"

	"github.com/ernanilima/gshopping/app/config"
	"github.com/ernanilima/gshopping/app/controller"
	brand_controller "github.com/ernanilima/gshopping/app/controller/brand"
	product_controller "github.com/ernanilima/gshopping/app/controller/product"
	"github.com/ernanilima/gshopping/app/repository"
	"github.com/ernanilima/gshopping/app/repository/database"
	"github.com/stretchr/testify/assert"
)

// Deve retornar o total de metodos existentes
func TestController_Should_Return_The_Total_Methods(t *testing.T) {
	totalMethods := reflect.TypeOf((*controller.Controller)(nil)).Elem().NumMethod()
	assert.Equal(t, 4, totalMethods)
}

// Deve retornar os metodos para a interface product controller
func TestNewController_Should_Return_Methods_For_ProductController(t *testing.T) {
	databaseConfig := &database.DatabaseConfig{Config: &config.Config{}}
	controller := controller.NewController(repository.NewRepository(databaseConfig))

	result, exist := controller.(product_controller.ProductController)
	assert.True(t, exist)
	productControllerTypeOf := reflect.TypeOf((*product_controller.ProductController)(nil)).Elem()
	controllerTypeOf := reflect.TypeOf(result)
	assert.NotNil(t, controllerTypeOf)

	nameOfMethods := []string{
		"FindProductByBarcode",
	}

	assert.Equal(t, len(nameOfMethods), productControllerTypeOf.NumMethod())
	for _, name := range nameOfMethods {
		// metodos existentes no controlador generico
		method, exist := controllerTypeOf.MethodByName(name)
		assert.True(t, exist)
		assert.Equal(t, name, method.Name)

		// metodos existentes no controlador especifico
		method, exist = productControllerTypeOf.MethodByName(name)
		assert.True(t, exist)
		assert.Equal(t, name, method.Name)
	}
}

// Deve retornar os metodos para a interface brand controller
func TestNewController_Should_Return_Methods_For_BrandController(t *testing.T) {
	databaseConfig := &database.DatabaseConfig{Config: &config.Config{}}
	controller := controller.NewController(repository.NewRepository(databaseConfig))

	result, exist := controller.(brand_controller.BrandController)
	assert.True(t, exist)
	brandControllerTypeOf := reflect.TypeOf((*brand_controller.BrandController)(nil)).Elem()
	controllerTypeOf := reflect.TypeOf(result)
	assert.NotNil(t, controllerTypeOf)

	nameOfMethods := []string{
		"FindAllBrands",
		"FindBrandById",
		"FindAllBrandsByDescription",
	}

	assert.Equal(t, len(nameOfMethods), brandControllerTypeOf.NumMethod())
	for _, name := range nameOfMethods {
		// metodos existentes no controlador generico
		method, exist := controllerTypeOf.MethodByName(name)
		assert.True(t, exist)
		assert.Equal(t, name, method.Name)

		// metodos existentes no controlador especifico
		method, exist = brandControllerTypeOf.MethodByName(name)
		assert.True(t, exist)
		assert.Equal(t, name, method.Name)
	}
}
