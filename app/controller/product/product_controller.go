package product_controller

import (
	"net/http"

	"github.com/ernanilima/gshopping/app/utils/response"
	"github.com/go-chi/chi"
)

// FindProductByBarcode busca um produto pelo codigo de barras
func (repo *productRepository) FindProductByBarcode(w http.ResponseWriter, r *http.Request) {

	product, err := repo.ProductRepository.FindByBarcode(chi.URLParam(r, "barcode"))
	if err != nil {
		messageError := "Produto não encontrado"
		response.Error(w, r, http.StatusNotFound, messageError)
		return
	}

	response.JSON(w, http.StatusOK, product)
}
