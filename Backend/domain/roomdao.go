package domain

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type RoomDAO struct {
	ID         pgtype.Int4
	RoomType   pgtype.Text
	BedType    pgtype.Text
	Rent       pgtype.Int4
	PropertyID pgtype.Int4
}

func (dao *RoomDAO) MapToDomain(dom *Room) {
	dom.ID = uint64(dao.ID.Int32)
	dom.BedType = dao.BedType.String
	dom.RoomType = dao.RoomType.String
	dom.PropertyID = uint64(dao.PropertyID.Int32)
	dom.Rent = uint64(dao.Rent.Int32)
}
