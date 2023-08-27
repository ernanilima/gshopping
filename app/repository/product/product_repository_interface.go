package product_repository

import (
	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/repository/database"
	"github.com/ernanilima/gshopping/app/utils"
	"github.com/google/uuid"
)

func NewProductRepository(connector database.DatabaseConnector) ProductRepository {
	return &ProductConnection{connector}
}

type ProductConnection struct {
	database.DatabaseConnector
}

type ProductRepository interface {
	InsertProduct(model.Product) (model.Product, error)
	EditProduct(model.Product) (model.Product, error)
	FindAllProducts(filter string, pageable utils.Pageable) utils.Pageable
	FindProductById(id uuid.UUID) (model.Product, error)
	FindProductByBarcode(barcode string) (model.Product, error)
	FindAllProductsNotFound(pageable utils.Pageable) utils.Pageable
	FindAllProductsNotFoundByBarcode(barcode string, pageable utils.Pageable) (utils.Pageable, error)
	FindTotalProducts() int32
	FindTotalProductsNotFound() int32
}
