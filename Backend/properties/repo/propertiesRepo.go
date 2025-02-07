package repo

import (
	"context"

	"github.com/Kenasvarghese/Booking-App/Backend/database"
	"github.com/Kenasvarghese/Booking-App/Backend/domain"
)

type propertiesRepo struct {
	db database.DB
}

func NewPropertiesRepo(db database.DB) domain.PropertiesRepo {
	return &propertiesRepo{
		db: db,
	}
}

var (
	queryListAllProperties = `
	SELECT 
		id,
		name,
		room_count
	FROM 
		properties
	`
)

func (r *propertiesRepo) ListAllProperties(ctx context.Context) ([]domain.Property, error) {
	query := queryListAllProperties
	args := make([]any, 0)
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return []domain.Property{}, err
	}
	defer rows.Close()
	properties := make([]domain.Property, 0)
	for rows.Next() {
		var property domain.Property
		var dao domain.PropertyDAO
		if err := rows.Scan(
			&dao.ID,
			&dao.Name,
			&dao.RoomCount,
		); err != nil {
			return []domain.Property{}, err
		}
		dao.MapToDomain(&property)
		properties = append(properties, property)
	}
	if rows.Err() != nil {
		return []domain.Property{}, err
	}
	return properties, nil
}
