package domain

import "context"

type PropertiesUsecase interface {
	ListAllProperties(ctx context.Context) ([]Property, error)
	AddProperty(ctx context.Context, property AddPropertyDTO) (uint64, error)
}

type Property struct {
	Name      string
	ID        uint64
	RoomCount uint64
	Address   string
}
