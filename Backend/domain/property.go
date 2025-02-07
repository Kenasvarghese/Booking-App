package domain

import "context"

type PropertiesUsecase interface {
	ListAllProperties(ctx context.Context) ([]Property, error)
}

type Property struct {
	Name      string
	ID        uint64
	RoomCount uint64
}
