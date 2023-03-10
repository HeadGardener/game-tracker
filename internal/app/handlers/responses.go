package handlers

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func newErrResponse(w http.ResponseWriter, code int, errorMsg string) {
	zap.L().Error(errorMsg)
	newResponse(w, code, response{
		Message: errorMsg,
	})
}

func newResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
