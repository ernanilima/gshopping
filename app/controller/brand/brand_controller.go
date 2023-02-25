package brand_controller

import (
	"fmt"
	"net/http"

	"github.com/ernanilima/gshopping/app/utils"
	"github.com/ernanilima/gshopping/app/utils/response"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// FindAllBrands busca uma lista com todas as marcas
func (repo *brandRepository) FindAllBrands(w http.ResponseWriter, r *http.Request) {
	pagination := utils.PaginationFilters(r)
	response.JSON(w, http.StatusOK, repo.BrandRepository.FindAll(pagination))
}

// FindBrandById busca uma marca pelo ID
func (repo *brandRepository) FindBrandById(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		messageError := "ID inválido"
		response.Error(w, r, http.StatusUnprocessableEntity, messageError)
		return
	}

	brand, err := repo.BrandRepository.FindById(id)
	if err != nil {
		messageError := "Marca não encontrada"
		response.Error(w, r, http.StatusNotFound, messageError)
		return
	}

	response.JSON(w, http.StatusOK, brand)
}

// FindAllBrandsByDescription busca uma lista de marcas pela descricao
func (repo *brandRepository) FindAllBrandsByDescription(w http.ResponseWriter, r *http.Request) {

	description := chi.URLParam(r, "description")
	pagination := utils.PaginationFilters(r)
	brands, err := repo.BrandRepository.FindByDescription(description, pagination)
	if err != nil {
		messageError := "Marca não encontrada"
		response.Error(w, r, http.StatusNotFound, messageError)
		return
	}

	if brands.TotalElements == 0 {
		messageError := fmt.Sprintf("Nenhuma Marca encontrada com '%s'", description)
		response.Error(w, r, http.StatusNotFound, messageError)
		return
	}

	response.JSON(w, http.StatusOK, brands)
}
