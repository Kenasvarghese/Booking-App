package domain

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type PropertyDAO struct {
	ID        pgtype.Int4
	Name      pgtype.Text
	RoomCount pgtype.Int4
	Address   pgtype.Text
}
type PropertiesRepo interface {
	ListAllProperties(ctx context.Context) ([]Property, error)
	AddProperty(ctx context.Context, property Property) (uint64, error)
}

func (dao *PropertyDAO) MapToDomain(domain *Property) {
	domain.ID = uint64(dao.ID.Int32)
	domain.Name = dao.Name.String
	domain.RoomCount = uint64(dao.RoomCount.Int32)
	domain.Address = dao.Address.String
}
