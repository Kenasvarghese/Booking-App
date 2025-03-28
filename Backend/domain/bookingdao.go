package domain

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type BookingsRepo interface {
	AddBooking(ctx context.Context, booking Booking) (string, error)
	GetBookings(ctx context.Context) ([]Booking, error)
}

type BookingDAO struct {
	ID         pgtype.UUID
	RoomID     pgtype.Int4
	PropertyID pgtype.Int4
	UserID     pgtype.Int4
	CheckIn    pgtype.Timestamp
	CheckOut   pgtype.Timestamp
	Status     pgtype.Text
}

func (dao *BookingDAO) MapToDomain(domain *Booking) {
	domain.ID = dao.ID.String()
	domain.RoomID = uint64(dao.RoomID.Int32)
	domain.UserID = uint64(dao.UserID.Int32)
	domain.CheckIn = dao.CheckIn.Time
	domain.CheckOut = dao.CheckOut.Time
	domain.Status = dao.Status.String
	domain.PropertyID = uint64(dao.PropertyID.Int32)
}
