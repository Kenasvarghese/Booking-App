package domain

import "context"

type RoomsRepo interface {
	GetRoomsByPropertyID(ctx context.Context, propertyIDs []uint64) ([]Room, error)
	AddRoom(ctx context.Context, room Room) (uint64, error)
}

type RoomsUsecase interface {
	GetRoomsByPropertyID(ctx context.Context, propertyIDs []uint64) ([]Room, error)
	AddRoom(ctx context.Context, room AddRoomDTO) (uint64, error)
}

type Room struct {
	ID         uint64
	RoomType   string
	BedType    string
	Rent       uint64
	PropertyID uint64
}
