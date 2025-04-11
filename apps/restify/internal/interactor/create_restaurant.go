package interactor

import (
	"context"
	"javifood-restify/internal/domain"
	"javifood-restify/internal/domain/entity"
	"javifood-restify/internal/domain/repository"
	usecase "javifood-restify/internal/domain/usecase"
)

type CreateRestaurantInteractor struct {
	restaurantRepository repository.RestaurantRepository
}

func NewCreateRestaurantInteractor(
	restaurantRepository repository.RestaurantRepository,
) *CreateRestaurantInteractor {
	return &CreateRestaurantInteractor{
		restaurantRepository: restaurantRepository,
	}
}

func (uc *CreateRestaurantInteractor) Execute(
	ctx context.Context,
	input usecase.CreateRestaurantUseCaseInputDto,
) error {
	dbRestaurant, _ := uc.restaurantRepository.FindByUserID(ctx, input.UserID)
	if dbRestaurant != nil {
		return domain.UserAlreadyHaveRestaurantError
	}
	restaurant, err := entity.NewRestaurant(
		"",
		input.UserID,
		input.Name,
		input.Address,
		input.OpenTime,
		input.CloseTime,
		input.CoordinateX,
		input.CoordinateY,
		input.WorkDays,
		nil,
		nil,
	)
	if err != nil {
		return err
	}
	return uc.restaurantRepository.Store(ctx, restaurant)
}
