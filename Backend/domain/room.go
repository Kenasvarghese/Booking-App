package domain

import "context"

type RoomsRepo interface {
	GetRoomsByPropertyID(ctx context.Context, propertyID uint64) ([]Room, error)
}

type RoomsUsecase interface {
	GetRoomsByPropertyID(ctx context.Context, propertyID uint64) ([]Room, error)
}

type Room struct {
	ID         uint64
	RoomType   string
	BedType    string
	Rent       uint64
	PropertyID uint64
}
