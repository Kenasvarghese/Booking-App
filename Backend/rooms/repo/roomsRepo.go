package repo

import (
	"context"
	"fmt"

	"github.com/Kenasvarghese/Booking-App/Backend/database"
	"github.com/Kenasvarghese/Booking-App/Backend/domain"
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
		property_id IN ($1)
	`
)

func (r *roomsRepo) GetRoomsByPropertyID(ctx context.Context, propertyID uint64) ([]domain.Room, error) {
	rows, err := r.db.Query(ctx, queryGetRoomsByPropertyID, propertyID)
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
