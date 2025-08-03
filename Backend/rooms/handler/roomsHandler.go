package handler

import (
	"context"
	"encoding/json"
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
	r.HandleFunc("/room", handler.AddRoom).Methods("POST")
}

func (h *roomsHandler) ListRooms(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("property_id")
	if id == "" {
		utils.ApiErrorResponse(w, http.StatusBadRequest, "property_id missing")
		return
	}
	propertIDs := make([]uint64, 0)
	propertyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusBadRequest, "property_id not an integer")
		return
	}
	propertIDs = append(propertIDs, propertyID)
	rooms, err := h.roomsUsecase.GetRoomsByPropertyID(context.Background(), propertIDs)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	roomsDto := make([]domain.RoomDTO, len(rooms))
	for _, room := range rooms {
		roomsDto = append(roomsDto, domain.RoomDTO(room))
	}
	utils.ApiSuccessResponse(w, roomsDto)
}

func (h *roomsHandler) AddRoom(w http.ResponseWriter, r *http.Request) {
	var room domain.AddRoomDTO
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = utils.Validate(room)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.roomsUsecase.AddRoom(context.Background(), room)
	if err != nil {
		utils.ApiErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.ApiSuccessResponse(w, struct {
		Id uint64 `json:"room_id"`
	}{
		Id: id,
	})
}
