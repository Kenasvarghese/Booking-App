package handler

import (
	"context"
	"net/http"

	"github.com/Kenasvarghese/Booking-App/Backend/domain"
	"github.com/Kenasvarghese/Booking-App/Backend/utils"
	"github.com/gorilla/mux"
)

type propertiesHandler struct {
	propertiesUsecase domain.PropertiesUsecase
}

func NewPropertiesHandler(r *mux.Router, propertiesUsecase domain.PropertiesUsecase) {
	handler := &propertiesHandler{
		propertiesUsecase: propertiesUsecase,
	}
	r.HandleFunc("/properties", handler.ListProperties).Methods("GET")
}

func (h *propertiesHandler) ListProperties(w http.ResponseWriter, r *http.Request) {
	properties, err := h.propertiesUsecase.ListAllProperties(context.Background())
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusInternalServerError, err)
	}
	utils.ApiSuccessResponse(w, properties)
}
