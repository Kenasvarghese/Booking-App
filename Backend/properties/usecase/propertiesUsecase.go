package usecase

import (
	"context"

	"github.com/Kenasvarghese/Booking-App/Backend/domain"
)

type propertiesUsecase struct {
	propertiesRepo domain.PropertiesRepo
}

func NewPropertiesUsecaseHandler(propertiesRepo domain.PropertiesRepo) domain.PropertiesUsecase {
	return &propertiesUsecase{
		propertiesRepo: propertiesRepo,
	}
}

// ListAllProperties usecase for listing all property
func (u *propertiesUsecase) ListAllProperties(ctx context.Context) ([]domain.Property, error) {
	return u.propertiesRepo.ListAllProperties(ctx)
}

// AddProperty usecase for adding new property
func (u *propertiesUsecase) AddProperty(ctx context.Context, property domain.AddPropertyDTO) (uint64, error) {
	return u.propertiesRepo.AddProperty(ctx, domain.Property{
		Name:      property.Name,
		RoomCount: property.RoomCount,
		Address:   property.Address,
	})
}
