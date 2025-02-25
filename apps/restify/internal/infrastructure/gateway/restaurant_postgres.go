package gateway

import (
	"context"
	"javifood-restify/internal/domain/entity"
	"javifood-restify/internal/infrastructure/database"
	"javifood-restify/internal/infrastructure/database/mapper"

	"gorm.io/gorm"
)

type RestaurantRepository struct {
	db               *gorm.DB
	restaurantMapper mapper.RestaurantMapper
}

func NewRestaurantRepository() *RestaurantRepository {
	return &RestaurantRepository{
		db:               database.DBConn,
		restaurantMapper: mapper.RestaurantMapper{},
	}
}

func (rp *RestaurantRepository) Store(ctx context.Context, restaurant *entity.Restaurant) error {
	restaurantDB, err := rp.restaurantMapper.ToDatabase(*restaurant)
	if err != nil {
		return err
	}
	tx := rp.db.Create(&restaurantDB)
	return tx.Error
}

func (rp *RestaurantRepository) FindByUserID(
	ctx context.Context,
	userID string,
) (*entity.Restaurant, error) {
	return nil, nil
}
