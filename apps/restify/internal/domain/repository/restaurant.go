package repository

import (
	"context"
	entity "javifood-restify/internal/domain/entity"
)

type RestaurantRepository interface {
    Store(ctx context.Context, restaurant *entity.Restaurant) error
    FindByUserID(ctx context.Context, userID string) (*entity.Restaurant, error)
}
