package domain

import (
    "time"
    entity "javifood-restify/internal/domain/entity"
)

type CreateRestaurantInputDto struct {
	UserID     string
	Name       string
	Address    string
	Coordinate *entity.Coordinate
	OpenTime   time.Time
	Closetime  time.Time
	WorkDays   *entity.WorkDays
}

type CreateRestaurant interface {
    handle(input CreateRestaurantInputDto)
}
