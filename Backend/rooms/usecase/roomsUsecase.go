package usecase

import (
	"context"

	"github.com/Kenasvarghese/Booking-App/Backend/domain"
)

type roomsUsecase struct {
	roomsRepo domain.RoomsRepo
}

func NewRoomUsecase(roomsRepo domain.RoomsRepo) domain.RoomsUsecase {
	return &roomsUsecase{
		roomsRepo: roomsRepo,
	}
}

func (u *roomsUsecase) GetRoomsByPropertyID(ctx context.Context, propertyID uint64) ([]domain.Room, error) {
	return u.roomsRepo.GetRoomsByPropertyID(ctx, propertyID)
}
