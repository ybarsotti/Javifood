package gatewaymock

import (
	"context"
	"javifood-restify/internal/domain/entity"
)

type RestaurantInMemoryRepository struct {
	restaurants []entity.Restaurant
}

func NewRestaurantInMemoryRepository() *RestaurantInMemoryRepository {
	return &RestaurantInMemoryRepository{}
}

func (r *RestaurantInMemoryRepository) Store(
	ctx context.Context,
	restaurant *entity.Restaurant,
) error {
	r.restaurants = append(r.restaurants, *restaurant)
	return nil
}

func (r *RestaurantInMemoryRepository) FindByUserID(
	ctx context.Context,
	userID string,
) (*entity.Restaurant, error) {
	for _, restaurant := range r.restaurants {
		if restaurant.UserID == userID {
			return &restaurant, nil
		}
	}
	return nil, nil
}
