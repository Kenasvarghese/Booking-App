package usecase

import (
	"context"
	"time"

	"github.com/Kenasvarghese/Booking-App/Backend/domain"
)

type bookingsUsecase struct {
	bookingsRepo domain.BookingsRepo
}

func NewBookingsUsecase(bookingsRepo domain.BookingsRepo) domain.BookingsUsecase {
	return &bookingsUsecase{
		bookingsRepo: bookingsRepo,
	}
}

func (u *bookingsUsecase) GetBookings(ctx context.Context) ([]domain.Booking, error) {
	return u.bookingsRepo.GetBookings(ctx)

}

// GetBookings retrieves all bookings from the repository
func (u *bookingsUsecase) AddBooking(ctx context.Context, bookingDTO domain.BookingDTO) (string, error) {
	checkIn, err := time.Parse("2006-01-02", bookingDTO.CheckIn)
	if err != nil {
		return "", err
	}
	checkOut, err := time.Parse("2006-01-02", bookingDTO.CheckOut)
	if err != nil {
		return "", err
	}

	return u.bookingsRepo.AddBooking(ctx, domain.Booking{
		RoomID:     bookingDTO.RoomID,
		UserID:     bookingDTO.UserID,
		CheckIn:    checkIn,
		CheckOut:   checkOut,
		Status:     bookingDTO.Status,
		PropertyID: bookingDTO.PropertyID,
	})
}
