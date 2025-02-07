package domain

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type PropertyDAO struct {
	ID        pgtype.Int4
	Name      pgtype.Text
	RoomCount pgtype.Int4
}
type PropertiesRepo interface {
	ListAllProperties(ctx context.Context) ([]Property, error)
}

func (dao *PropertyDAO) MapToDomain(dto *Property) {
	dto.ID = uint64(dao.ID.Int32)
	dto.Name = dao.Name.String
	dto.RoomCount = uint64(dao.RoomCount.Int32)
}
