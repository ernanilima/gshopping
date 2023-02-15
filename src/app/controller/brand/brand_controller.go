package brand

import (
	"encoding/json"
	"net/http"

	"github.com/ernanilima/gshopping/src/app/repository/brand"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// FindById busca uma marca pelo ID
func FindById(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.DefaultMaxHeaderBytes)
		return
	}

	brand, err := brand.FindById(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.DefaultMaxHeaderBytes)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brand)
}
