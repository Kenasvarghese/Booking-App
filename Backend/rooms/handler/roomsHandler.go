package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Kenasvarghese/Booking-App/Backend/domain"
	"github.com/Kenasvarghese/Booking-App/Backend/utils"
	"github.com/gorilla/mux"
)

type roomsHandler struct {
	roomsUsecase domain.RoomsUsecase
}

func NewRoomsHandler(r *mux.Router, roomsUsecase domain.RoomsUsecase) {
	handler := &roomsHandler{
		roomsUsecase: roomsUsecase,
	}
	r.HandleFunc("/rooms", handler.ListRooms).Methods("GET")
}

func (h *roomsHandler) ListRooms(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("property-id")
	if id == "" {
		utils.ApiErrorResponse(w, http.StatusBadRequest, "property-id missing", nil)
		return
	}
	propertyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusBadRequest, "property-id not an integer", err)
		return
	}
	rooms, err := h.roomsUsecase.GetRoomsByPropertyID(context.Background(), propertyID)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusInternalServerError, "", err)
		return
	}
	roomsDto := make([]domain.RoomDTO, len(rooms))
	for _, room := range rooms {
		roomsDto = append(roomsDto, domain.RoomDTO(room))
	}
	utils.ApiSuccessResponse(w, roomsDto)
}
