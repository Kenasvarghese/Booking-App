package utils

import (
	"encoding/json"
	"net/http"
)

func ApiSuccessResponse(w http.ResponseWriter, data any) {
	response := struct {
		Success bool `json:"success"`
		Data    any  `json:"data,omitempty"`
	}{
		Success: true,
		Data:    data,
	}
	databyte, err := json.Marshal(response)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(databyte)
}

func ApiErrorResponse(w http.ResponseWriter, status int, err error) {
	response := struct {
		Success bool `json:"success"`
		Error   any  `json:"error"`
	}{
		Success: false,
		Error:   err,
	}
	databyte, err := json.Marshal(response)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(databyte)
}
