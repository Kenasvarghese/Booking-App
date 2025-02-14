package handler

import (
	"context"
	"encoding/json"
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
	r.HandleFunc("/property", handler.AddProperty).Methods("POST")
}

// ListProperties proeperties listing api handler
func (h *propertiesHandler) ListProperties(w http.ResponseWriter, r *http.Request) {
	properties, err := h.propertiesUsecase.ListAllProperties(context.Background())
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusInternalServerError, "", err)
	}
	propertiesDto := make([]domain.PropertyDTO, 0)
	for _, property := range properties {
		var dto domain.PropertyDTO
		dto.MapFromDomain(property)
		propertiesDto = append(propertiesDto, dto)
	}

	utils.ApiSuccessResponse(w, propertiesDto)
}

// AddProperty handler for adding new property
func (h *propertiesHandler) AddProperty(w http.ResponseWriter, r *http.Request) {
	var addPropertyDto domain.AddPropertyDTO
	err := json.NewDecoder(r.Body).Decode(&addPropertyDto)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusBadRequest, "", err)
	}
	err = utils.Validate(addPropertyDto)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusBadRequest, "", err)
	}

	id, err := h.propertiesUsecase.AddProperty(context.Background(), addPropertyDto)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusInternalServerError, "", err)
	}
	utils.ApiSuccessResponse(w, struct {
		Id uint64 `json:"id"`
	}{
		Id: id,
	})
}
