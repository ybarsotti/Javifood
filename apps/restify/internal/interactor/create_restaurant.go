package interactor

import (
	"context"
	"errors"
	"javifood-restify/internal/domain"
	"javifood-restify/internal/domain/entity"
	"javifood-restify/internal/domain/repository"
	usecase "javifood-restify/internal/domain/usecase"
)

type CreateRestaurantInteractor struct {
	restaurantRepository repository.RestaurantRepository
}

func NewCreateRestaurantInteractor(restaurantRepository repository.RestaurantRepository) *CreateRestaurantInteractor {
	return &CreateRestaurantInteractor{
		restaurantRepository: restaurantRepository,
	}
}

// CreateRestaurant godoc
//
//	@Summary	Create a restaurant
//	@Tags		restaurant
//	@Accept		json
//	@Success	201
//	@Router		/api/v1/restaurants [post]
func (uc *CreateRestaurantInteractor) Execute(ctx context.Context, input usecase.CreateRestaurantUseCaseInputDto) error {
	dbRestaurant, err := uc.restaurantRepository.FindByUserID(ctx, input.UserID)
	if dbRestaurant != nil {
		return domain.UserAlreadyHaveRestaurantError
	}
	if err != nil {
		return errors.Join(domain.InternalServerError, err)
	}
	restaurant, err := entity.NewRestaurant(
		"", input.UserID, input.Name, input.Address, input.CoordinateX, input.CoordinateY, input.OpenTimeHour, input.OpenTimeMinute, input.CloseTimeHour, input.CloseTimeMinute, input.WorkDays, nil, nil,
	)
	return uc.restaurantRepository.Store(ctx, restaurant)
}
