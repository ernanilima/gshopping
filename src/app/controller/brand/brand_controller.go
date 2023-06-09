package brand

import (
	"fmt"
	"net/http"

	"github.com/ernanilima/gshopping/src/app/repository/brand"
	"github.com/ernanilima/gshopping/src/app/utils"
	"github.com/ernanilima/gshopping/src/app/utils/response"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// FindAll busca uma lista com todas as marcas
func FindAll(w http.ResponseWriter, r *http.Request) {
	pagination := utils.PaginationFilters(r)
	response.JSON(w, http.StatusOK, brand.FindAll(pagination))
}

// FindById busca uma marca pelo ID
func FindById(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		messageError := "ID inválido"
		response.Error(w, r, http.StatusUnprocessableEntity, messageError)
		return
	}

	brand, err := brand.FindById(id)
	if err != nil {
		messageError := "Marca não encontrada"
		response.Error(w, r, http.StatusNotFound, messageError)
		return
	}

	response.JSON(w, http.StatusOK, brand)
}

// FindByDescription busca uma lista de marcas pela descricao
func FindByDescription(w http.ResponseWriter, r *http.Request) {

	description := chi.URLParam(r, "description")
	pagination := utils.PaginationFilters(r)
	brands, err := brand.FindByDescription(description, pagination)
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
