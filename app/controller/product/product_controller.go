package product_controller

import (
	"fmt"
	"net/http"

	"github.com/ernanilima/gshopping/app/utils"
	"github.com/ernanilima/gshopping/app/utils/response"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// FindAllProducts busca uma lista com todos os produtos
func (repo *productRepository) FindAllProducts(w http.ResponseWriter, r *http.Request) {

	filter := chi.URLParam(r, "filter")
	pagination := utils.PaginationFilters(r)
	products := repo.ProductRepository.FindAllProducts(filter, pagination)
	if products.TotalElements == 0 && len(filter) > 0 {
		messageError := fmt.Sprintf("Nenhum Produto encontrado com '%s'", filter)
		response.Error(w, r, http.StatusNotFound, messageError)
		return
	}

	response.JSON(w, http.StatusOK, products)
}

// FindProductById busca um produto pelo ID
func (repo *productRepository) FindProductById(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		messageError := "ID inválido"
		response.Error(w, r, http.StatusUnprocessableEntity, messageError)
		return
	}

	brand, err := repo.ProductRepository.FindProductById(id)
	if err != nil {
		messageError := "Produto não encontrado"
		response.Error(w, r, http.StatusNotFound, messageError)
		return
	}

	response.JSON(w, http.StatusOK, brand)
}

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

// FindAllProductsNotFound busca uma lista com todos os produtos nao encontrados
func (repo *productRepository) FindAllProductsNotFound(w http.ResponseWriter, r *http.Request) {
	pagination := utils.PaginationFilters(r)
	response.JSON(w, http.StatusOK, repo.FindAllNotFound(pagination))
}

// FindAllProductsNotFoundByBarcode busca uma lista de produtos nao encontrados pelo codigo de barras
func (repo *productRepository) FindAllProductsNotFoundByBarcode(w http.ResponseWriter, r *http.Request) {

	barcode := chi.URLParam(r, "barcode")
	pagination := utils.PaginationFilters(r)
	productsNotFound, err := repo.ProductRepository.FindNotFoundByBarcode(barcode, pagination)
	if err != nil || productsNotFound.TotalElements == 0 {
		messageError := fmt.Sprintf("Nenhum registro encontrado com '%s'", barcode)
		response.Error(w, r, http.StatusNotFound, messageError)
		return
	}

	response.JSON(w, http.StatusOK, productsNotFound)
}
