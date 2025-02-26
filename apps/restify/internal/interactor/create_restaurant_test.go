package interactor_test

import (
	"context"
	"errors"
	"javifood-restify/internal/domain"
	"javifood-restify/internal/domain/usecase"
	gatewaymock "javifood-restify/internal/infrastructure/gateway/gateway_mock"

	"javifood-restify/internal/interactor"
	"testing"

	"github.com/google/uuid"
)

func newUseCase(
) *interactor.CreateRestaurantInteractor {
	return interactor.NewCreateRestaurantInteractor(
		gatewaymock.NewRestaurantInMemoryRepository(),
	)
}

func TestCreateRestaurant_CreateSuccessfully(t *testing.T) {
	createRestaurant := newUseCase()
	ctx := context.Background()
	input := usecase.NewCreateRestaurantInputDto(
		uuid.NewString(),
		"Restaurant 1",
		"Address",
		21.222,
		-10.000,
		10,
		00,
		20,
		00,
		[]string{"Monday"},
	)
	err := createRestaurant.Execute(ctx, *input)
	if err != nil {
		t.Errorf("should not have error")
	}
}

func TestCreateRestaurant_ConflictError(t *testing.T) {
	createRestaurant := newUseCase()
	ctx := context.Background()
	input := usecase.NewCreateRestaurantInputDto(
		uuid.NewString(),
		"Restaurant 1",
		"Address",
		21.222,
		-10.000,
		10,
		00,
		20,
		00,
		[]string{"Monday"},
	)
	err := createRestaurant.Execute(ctx, *input)
	err = createRestaurant.Execute(ctx, *input)
	if !errors.Is(err, domain.UserAlreadyHaveRestaurantError) {
		t.Errorf("should raise UserAlreadyHaveRestaurantError")
	}
}

