package domain

type BookingDTO struct {
	ID         string `json:"id"`
	RoomID     uint64 `json:"room_id"`
	UserID     uint64 `json:"user_id"`
	Email      string `json:"email"`
	CheckIn    string `json:"check_in"`
	CheckOut   string `json:"check_out"`
	Status     string `json:"status"`
	PropertyID uint64 `json:"property_id"`
}
