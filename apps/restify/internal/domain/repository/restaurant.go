package repository

import (
	"context"
	entity "javifood-restify/internal/domain/entity"

	"github.com/google/uuid"
)

type RestaurantRepository interface {
	Store(ctx context.Context, restaurant *entity.Restaurant) error
	FindByUserID(ctx context.Context, userID uuid.UUID) (*entity.Restaurant, error)
}
