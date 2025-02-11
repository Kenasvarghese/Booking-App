package domain

type PropertyDTO struct {
	Name      string `json:"name"`
	ID        uint64 `json:"id"`
	RoomCount uint64 `json:"room_count"`
	Address   string `json:"address"`
}

func (dto *PropertyDTO) MapFromDomain(domain Property) {
	dto.ID = domain.ID
	dto.Name = domain.Name
	dto.RoomCount = domain.RoomCount
	dto.Address = domain.Address
}

type AddPropertyDTO struct {
	Name      string `json:"name" validate:"required"`
	RoomCount uint64 `json:"room_count" validate:"required"`
	Address   string `json:"address" validate:"required"`
}
