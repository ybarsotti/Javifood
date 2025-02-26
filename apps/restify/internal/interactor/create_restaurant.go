package interactor

import (
	"context"
	"errors"
	"javifood-restify/internal/domain"
	"javifood-restify/internal/domain/entity"
	"javifood-restify/internal/domain/repository"
	usecase "javifood-restify/internal/domain/usecase"

	"github.com/google/uuid"
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
	userUUID, err := uuid.Parse(input.UserID)
	if err != nil {
		return err
	}
	dbRestaurant, err := uc.restaurantRepository.FindByUserID(ctx, userUUID)
	if dbRestaurant != nil {
		return domain.UserAlreadyHaveRestaurantError
	}
	if err != nil {
		return errors.Join(domain.InternalServerError, err)
	}
	restaurant, err := entity.NewRestaurant(
		"",
		input.UserID,
		input.Name,
		input.Address,
		input.CoordinateX,
		input.CoordinateY,
		input.OpenTimeHour,
		input.OpenTimeMinute,
		input.CloseTimeHour,
		input.CloseTimeMinute,
		input.WorkDays,
		nil,
		nil,
	)
	if err != nil {
		return err
	}
	return uc.restaurantRepository.Store(ctx, restaurant)
}
