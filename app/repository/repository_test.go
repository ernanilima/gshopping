package repository_test

import (
	"reflect"
	"testing"

	"github.com/ernanilima/gshopping/app/config"
	"github.com/ernanilima/gshopping/app/repository"
	brand_repository "github.com/ernanilima/gshopping/app/repository/brand"
	"github.com/ernanilima/gshopping/app/repository/database"
	product_repository "github.com/ernanilima/gshopping/app/repository/product"
	"github.com/stretchr/testify/assert"
)

// Deve retornar o total de metodos existentes
func TestRepository_Should_Return_The_Total_Methods(t *testing.T) {
	totalMethods := reflect.TypeOf((*repository.Repository)(nil)).Elem().NumMethod()
	assert.Equal(t, 13, totalMethods)
}

// Deve retornar os metodos para a interface product repository
func TestNewRepository_Should_Return_Methods_For_ProductRepository(t *testing.T) {
	databaseConfig := &database.DatabaseConfig{Config: &config.Config{}}
	repository := repository.NewRepository(databaseConfig)

	result, exist := repository.(product_repository.ProductRepository)
	assert.True(t, exist)
	productRepositoryTypeOf := reflect.TypeOf((*product_repository.ProductRepository)(nil)).Elem()
	repositoryTypeOf := reflect.TypeOf(result)
	assert.NotNil(t, repositoryTypeOf)

	nameOfMethods := []string{
		"InsertProduct",
		"EditProduct",
		"FindAllProducts",
		"FindProductById",
		"FindProductByBarcode",
		"FindAllProductsNotFound",
		"FindAllProductsNotFoundByBarcode",
	}

	assert.Equal(t, len(nameOfMethods), productRepositoryTypeOf.NumMethod())
	for _, name := range nameOfMethods {
		// metodos existentes no repositorio generico
		method, exist := repositoryTypeOf.MethodByName(name)
		assert.True(t, exist)
		assert.Equal(t, name, method.Name)

		// metodos existentes no repositorio especifico
		method, exist = productRepositoryTypeOf.MethodByName(name)
		assert.True(t, exist)
		assert.Equal(t, name, method.Name)
	}
}

// Deve retornar os metodos para a interface brand repository
func TestNewRepository_Should_Return_Methods_For_BrandRepository(t *testing.T) {
	databaseConfig := &database.DatabaseConfig{Config: &config.Config{}}
	repository := repository.NewRepository(databaseConfig)

	result, exist := repository.(brand_repository.BrandRepository)
	assert.True(t, exist)
	brandRepositoryTypeOf := reflect.TypeOf((*brand_repository.BrandRepository)(nil)).Elem()
	repositoryTypeOf := reflect.TypeOf(result)
	assert.NotNil(t, repositoryTypeOf)

	nameOfMethods := []string{
		"InsertBrand",
		"EditBrand",
		"DeleteBrand",
		"FindAllBrands",
		"FindBrandById",
		"FindAllBrandsByDescription",
	}

	assert.Equal(t, len(nameOfMethods), brandRepositoryTypeOf.NumMethod())
	for _, name := range nameOfMethods {
		// metodos existentes no repositorio generico
		method, exist := repositoryTypeOf.MethodByName(name)
		assert.True(t, exist)
		assert.Equal(t, name, method.Name)

		// metodos existentes no repositorio especifico
		method, exist = brandRepositoryTypeOf.MethodByName(name)
		assert.True(t, exist)
		assert.Equal(t, name, method.Name)
	}
}
