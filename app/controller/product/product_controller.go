package product_controller

import "net/http"

// FindByBarcode busca a descricao
func FindByBarcode(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um produto pelo codigo de barras"))
}
