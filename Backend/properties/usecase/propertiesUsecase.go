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
func (u *propertiesUsecase) ListAllProperties(ctx context.Context) ([]domain.Property, error) {
	return u.propertiesRepo.ListAllProperties(ctx)

}
