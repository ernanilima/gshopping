package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ernanilima/gshopping/app/utils"
)

type StandardError struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

// JSON eh usado para retornar sucesso
func JSONPageable(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	dataPageable := utils.Pageable{
		Content:       data,
		TotalPages:    1,
		TotalElements: 10,
		Size:          111,
		Page:          0,
	}

	json.NewEncoder(w).Encode(dataPageable)
}

// JSON eh usado para retornar sucesso
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Error eh usado para retornar um erro personalizado
func Error(w http.ResponseWriter, r *http.Request, statusCode int, msg string) {
	err := StandardError{
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    statusCode,
		Error:     http.StatusText(statusCode),
		Message:   msg,
		Path:      r.URL.Path,
	}
	JSON(w, statusCode, err)
}
