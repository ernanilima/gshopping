package response

import (
	"encoding/json"
	"net/http"
	"time"
)

type standardError struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	Path      string `json:"path"`
}

// JSON eh usado para retornar sucesso
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// Error eh usado para retornar um erro personalizado
func Error(w http.ResponseWriter, r *http.Request, statusCode int, msg string) {
	err := standardError{
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    404,
		Error:     http.StatusText(statusCode),
		Message:   msg,
		Path:      r.URL.Path,
	}
	JSON(w, statusCode, err)
}
