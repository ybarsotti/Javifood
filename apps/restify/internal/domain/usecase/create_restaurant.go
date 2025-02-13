package domain

import (
	repository "javifood-restify/internal/domain/repository"
	valueobject "javifood-restify/internal/domain/value_object"
	"time"
)

type CreateRestaurantUseCaseInputDto struct {
	UserID     string
	Name       string
	Address    string
	Coordinate *valueobject.Coordinate
	OpenTime   time.Time
	Closetime  time.Time
	WorkDays   *valueobject.WorkDays
}

type CreateRestaurantUseCase interface {
	repository.RestaurantRepository
	Handle(input CreateRestaurantUseCaseInputDto)
}
