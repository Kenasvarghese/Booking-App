package domain

type RoomDTO struct {
	ID         uint64 `json:"id"`
	RoomType   string `json:"room_type"`
	BedType    string `json:"bed_type"`
	Rent       uint64 `json:"rent"`
	PropertyID uint64 `json:"property_id"`
}

type AddRoomDTO struct {
	RoomType   string `json:"room_type" validate:"required"`
	BedType    string `json:"bed_type"`
	Rent       uint64 `json:"rent" validate:"required"`
	PropertyID uint64 `json:"property_id" validate:"required"`
}

func (dto *RoomDTO) MapFromDomain(dom *Room) {
	dto.ID = dom.ID
	dto.BedType = dom.BedType
	dto.RoomType = dom.RoomType
	dto.Rent = dom.Rent
	dto.PropertyID = dom.PropertyID
}
