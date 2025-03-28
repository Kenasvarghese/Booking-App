package repo

import (
	"context"

	"github.com/Kenasvarghese/Booking-App/Backend/database"
	"github.com/Kenasvarghese/Booking-App/Backend/domain"
	"github.com/jackc/pgx/v5/pgtype"
)

type bookingsRepo struct {
	db database.DB
}

func NewBookingRepo(db database.DB) domain.BookingsRepo {
	return &bookingsRepo{
		db: db,
	}
}

const (
	queryAddBooking = `
	INSERT INTO bookings
	(room_id,user_id,check_in,check_out,status)
	VALUES($1,$2,$3,$4,$5)
	RETURNING id
	`
	queryGetBookings = `
	SELECT b.id, room_id, user_id, check_in, check_out, status, r.property_id
	FROM bookings b
	LEFT JOIN rooms r ON b.room_id = r.id
	`
)

// GetBookings retrieves all bookings from the database
func (r *bookingsRepo) GetBookings(ctx context.Context) ([]domain.Booking, error) {
	query := queryGetBookings
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []domain.Booking
	for rows.Next() {
		var booking domain.Booking
		var bookingDao domain.BookingDAO
		err := rows.Scan(&bookingDao.ID,
			&bookingDao.RoomID,
			&bookingDao.UserID,
			&bookingDao.CheckIn,
			&bookingDao.CheckOut,
			&bookingDao.Status,
			&bookingDao.PropertyID)
		if err != nil {
			return nil, err
		}
		bookingDao.MapToDomain(&booking)
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

// AddBooking adds a new booking to the database
func (r *bookingsRepo) AddBooking(ctx context.Context, booking domain.Booking) (string, error) {
	query := queryAddBooking
	args := []any{
		booking.RoomID,
		booking.UserID,
		booking.CheckIn,
		booking.CheckOut,
		booking.Status,
	}
	var id pgtype.UUID
	err := r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
