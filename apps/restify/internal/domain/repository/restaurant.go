package repository

import (
	"context"
	"github.com/google/uuid"
	entity "javifood-restify/internal/domain/entity"
)

type RestaurantRepository interface {
	Store(ctx context.Context, restaurant *entity.Restaurant) error
	FindByUserID(ctx context.Context, userID uuid.UUID) (*entity.Restaurant, error)
}
