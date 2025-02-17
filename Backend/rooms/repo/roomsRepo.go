package repo

import (
	"context"
	"fmt"

	"github.com/Kenasvarghese/Booking-App/Backend/database"
	"github.com/Kenasvarghese/Booking-App/Backend/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

type roomsRepo struct {
	db database.DB
}

func NewRoomsRepo(db database.DB) domain.RoomsRepo {
	return &roomsRepo{
		db: db,
	}
}

const (
	queryAddRoom = `
	INSERT INTO rooms(
		room_type,
		bed_type,
		rent,
		property_id
	)
	VALUES($1, $2, $3, $4)
	RETURNING id
	`

	queryGetRoomsByPropertyID = `
	SELECT 
		id,
		room_type,
		bed_type,
		rent,
		property_id
	FROM
		rooms
	WHERE 
		property_id = ANY($1)
	`
)

func (r *roomsRepo) GetRoomsByPropertyID(ctx context.Context, propertyIDs []uint64) ([]domain.Room, error) {
	rows, err := r.db.Query(ctx, queryGetRoomsByPropertyID, propertyIDs)
	if err != nil {
		return []domain.Room{}, fmt.Errorf("GetRoomsByPropertyID - Query returned with err %w", err)
	}
	rooms := make([]domain.Room, 0)
	for rows.Next() {
		var room domain.Room
		var roomDao domain.RoomDAO
		rows.Scan(
			&roomDao.ID,
			&roomDao.RoomType,
			&roomDao.BedType,
			&roomDao.Rent,
			&roomDao.PropertyID,
		)
		roomDao.MapToDomain(&room)
		rooms = append(rooms, room)
	}
	if rows.Err() != nil {
		return []domain.Room{}, fmt.Errorf("GetRoomsByPropertyID - rows returned with err %w", rows.Err())
	}
	return rooms, nil
}

func (r *roomsRepo) AddRoom(ctx context.Context, room domain.Room) (uint64, error) {
	var pgid pgtype.Int4
	err := r.db.QueryRow(ctx, queryAddRoom, room.RoomType, room.BedType, room.Rent, room.PropertyID).Scan(&pgid)
	if err != nil {
		return 0, fmt.Errorf("AddRoom - QueryRow returned with err %w", err)
	}
	return uint64(pgid.Int32), nil
}
