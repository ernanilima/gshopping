package product_repository

import (
	"github.com/ernanilima/gshopping/app/repository/database"
)

func NewProductRepository(connector database.DatabaseConnector) ProductRepository {
	return &ProductConnection{connector}
}

type ProductConnection struct {
	database.DatabaseConnector
}

type ProductRepository interface {
	FindByBarcode(barcode string) (interface{}, error)
}
