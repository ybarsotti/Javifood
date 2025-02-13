package interactor

import (
	"context"
	"errors"
	"javifood-restify/internal/domain"
	"javifood-restify/internal/domain/repository"
	usecase "javifood-restify/internal/domain/usecase"
)

type CreateRestaurantInteractor struct {
	restaurantRepository repository.RestaurantRepository
}

func (uc *CreateRestaurantInteractor) Handle(ctx context.Context, input usecase.CreateRestaurantUseCaseInputDto) error {
    dbRestaurant, err := uc.restaurantRepository.FindByUserID(ctx, input.UserID)
    if dbRestaurant != nil {
        return domain.UserAlreadyHaveRestaurantError
    }
    if err != nil {
        return errors.Join(domain.InternalServerError, err)
    }
    return nil
}
