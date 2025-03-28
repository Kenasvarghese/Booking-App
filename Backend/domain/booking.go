package domain

import (
	"context"
	"time"
)

type BookingsUsecase interface {
	AddBooking(ctx context.Context, bookingDTO BookingDTO) (string, error)
	GetBookings(ctx context.Context) ([]Booking, error)
}

type Booking struct {
	ID        string
	RoomID     uint64
	UserID     uint64
	CheckIn    time.Time
	CheckOut   time.Time
	Status     string
	PropertyID uint64
}
