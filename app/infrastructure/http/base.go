package http

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   error       `json:"error"`
}

func JSONResponse(w http.ResponseWriter, httpCode int, success bool, message string, data interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(ResponseData{
		Success: success,
		Message: message,
		Data:    data,
		Error:   err,
	})
}
