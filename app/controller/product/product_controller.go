package product_controller

import (
	"net/http"

	"github.com/ernanilima/gshopping/app/utils/response"
	"github.com/go-chi/chi"
)

func (repo *productRepository) FindProductByBarcode(w http.ResponseWriter, r *http.Request) {
	barcode := chi.URLParam(r, "barcode")
	result, _ := repo.ProductRepository.FindByBarcode(barcode)
	response.JSON(w, http.StatusOK, result)
}
