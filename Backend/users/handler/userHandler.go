package handler

import (
	"net/http"

	"github.com/Kenasvarghese/Booking-App/Backend/utils"
	"github.com/gorilla/mux"
)

type userHandler struct {
}

func NewUserHandler(r *mux.Router) {
	handler := &userHandler{}
	r.HandleFunc("/users", handler.ListUsers).Methods("GET")
}

func (h userHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	println("reached")
	utils.ApiSuccessResponse(w, "successfully reached")
}
