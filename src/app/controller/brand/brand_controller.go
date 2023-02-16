package brand

import (
	"net/http"

	"github.com/ernanilima/gshopping/src/app/repository/brand"
	"github.com/ernanilima/gshopping/src/app/utils/response"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

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
