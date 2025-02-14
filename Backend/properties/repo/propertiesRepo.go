package repo

import (
	"context"
	"fmt"

	"github.com/Kenasvarghese/Booking-App/Backend/database"
	"github.com/Kenasvarghese/Booking-App/Backend/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

type propertiesRepo struct {
	db database.DB
}

func NewPropertiesRepo(db database.DB) domain.PropertiesRepo {
	return &propertiesRepo{
		db: db,
	}
}

const (
	queryAddProperty = `
	INSERT INTO properties
	(name,room_count,address)
	VALUES($1,$2,$3)
	RETURNING id
	`
	queryListAllProperties = `
	SELECT 
		id,
		name,
		room_count,
		address
	FROM 
		properties
	`
)

// ListAllProperties lists all properties
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
			&dao.Address,
		); err != nil {
			return []domain.Property{}, fmt.Errorf("ListAllProperties - rows scan returned with err %w", err)
		}
		dao.MapToDomain(&property)
		properties = append(properties, property)
	}
	if rows.Err() != nil {
		return []domain.Property{}, fmt.Errorf("ListAllProperties - rows returned with err %w", rows.Err())
	}
	return properties, nil
}

// AddProperty adds a new property
func (r *propertiesRepo) AddProperty(ctx context.Context, property domain.Property) (uint64, error) {
	query := queryAddProperty
	var id pgtype.Int4

	err := r.db.QueryRow(ctx, query, property.Name, property.RoomCount, property.Address).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("AddProperty QueryRow returned with err %w", err)
	}
	return uint64(id.Int32), nil
}
