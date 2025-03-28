package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Kenasvarghese/Booking-App/Backend/domain"
	"github.com/Kenasvarghese/Booking-App/Backend/utils"
	"github.com/gorilla/mux"
)

type bookingHandler struct {
	bookingsUsecase domain.BookingsUsecase
}

func NewBookingHandler(r *mux.Router, bookingsUsecase domain.BookingsUsecase) {
	handler := &bookingHandler{
		bookingsUsecase: bookingsUsecase,
	}
	r.HandleFunc("/book", handler.AddBooking).Methods("POST")
	r.HandleFunc("/bookings", handler.GetBookings).Methods("GET")
}

func (h *bookingHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	bookings, err := h.bookingsUsecase.GetBookings(r.Context())
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ApiSuccessResponse(w, bookings)
}

func (h *bookingHandler) AddBooking(w http.ResponseWriter, r *http.Request) {
	var bookingDTO domain.BookingDTO
	err := json.NewDecoder(r.Body).Decode(&bookingDTO)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = utils.Validate(bookingDTO)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusBadRequest, "validation error")
		return
	}
	id, err := h.bookingsUsecase.AddBooking(r.Context(), bookingDTO)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ApiSuccessResponse(w, struct {
		ID string `json:"booking_id"`
	}{
		ID: id,
	})
}
