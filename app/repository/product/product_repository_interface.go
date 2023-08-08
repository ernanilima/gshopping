package product_repository

import (
	"github.com/ernanilima/gshopping/app/model"
	"github.com/ernanilima/gshopping/app/repository/database"
	"github.com/ernanilima/gshopping/app/utils"
)

func NewProductRepository(connector database.DatabaseConnector) ProductRepository {
	return &ProductConnection{connector}
}

type ProductConnection struct {
	database.DatabaseConnector
}

type ProductRepository interface {
	FindByBarcode(barcode string) (model.Product, error)
	FindAllNotFound(pageable utils.Pageable) utils.Pageable
	FindNotFoundByBarcode(barcode string, pageable utils.Pageable) (utils.Pageable, error)
}
