package gateway

import (
	"context"
	"errors"
	"javifood-restify/internal/domain/entity"
	"javifood-restify/internal/infrastructure/database"
	"javifood-restify/internal/infrastructure/database/mapper"
	"javifood-restify/internal/infrastructure/database/model"

	"github.com/google/uuid"
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
	restaurantDB := rp.restaurantMapper.ToDatabase(*restaurant)
	tx := rp.db.Create(&restaurantDB)
	return tx.Error
}

func (rp *RestaurantRepository) FindByUserID(
	ctx context.Context,
	userID string,
) (*entity.Restaurant, error) {
	dbRestaurant := model.Restaurant{}
	rp.db.Where("user_id = ?", userID).First(&dbRestaurant)
	if dbRestaurant.ID == uuid.Nil {
		return nil, nil
	}
	restaurant, err := rp.restaurantMapper.ToDomain(dbRestaurant)
	if err != nil {
		return nil, errors.New("failed to hydrate restaurant")
	}
	return restaurant, nil
}
