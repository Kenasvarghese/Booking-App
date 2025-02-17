package usecase

import (
	"context"

	"github.com/Kenasvarghese/Booking-App/Backend/domain"
)

type roomsUsecase struct {
	roomsRepo domain.RoomsRepo
}

func NewRoomUsecase(roomsRepo domain.RoomsRepo, ) domain.RoomsUsecase {
	return &roomsUsecase{
		roomsRepo: roomsRepo,
	}
}

func (u *roomsUsecase) GetRoomsByPropertyID(ctx context.Context, propertyIDs []uint64) ([]domain.Room, error) {
	return u.roomsRepo.GetRoomsByPropertyID(ctx, propertyIDs)
}

func (u *roomsUsecase) AddRoom(ctx context.Context, room domain.AddRoomDTO) (uint64, error) {
	return u.roomsRepo.AddRoom(ctx, domain.Room{
		RoomType:   room.RoomType,
		BedType:    room.BedType,
		Rent:       room.Rent,
		PropertyID: room.PropertyID,
	})
}
