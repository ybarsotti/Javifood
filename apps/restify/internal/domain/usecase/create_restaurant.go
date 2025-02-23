package usecase

import (
	"context"
)

type CreateRestaurantUseCaseInputDto struct {
	UserID          string
	Name            string
	Address         string
	CoordinateX     float64
	CoordinateY     float64
	OpenTimeHour    uint8
	OpenTimeMinute  uint8
	CloseTimeHour   uint8
	CloseTimeMinute uint8
	WorkDays        []string
}

func NewCreateRestaurantInputDto(userId, name, address string, coordinateX, coordinateY float64, openTimeHour, openTimeMinute, closeTimeHour, closeTimeMinute uint8, workDays []string) *CreateRestaurantUseCaseInputDto {
	return &CreateRestaurantUseCaseInputDto{
		UserID:          userId,
		Name:            name,
		Address:         address,
		CoordinateX:     coordinateX,
		CoordinateY:     coordinateY,
		OpenTimeHour:    openTimeHour,
		OpenTimeMinute:  openTimeMinute,
		CloseTimeHour:   closeTimeHour,
		CloseTimeMinute: closeTimeMinute,
		WorkDays:        workDays,
	}
}

type CreateRestaurantUseCase interface {
	Execute(ctx context.Context, input CreateRestaurantUseCaseInputDto) error
}
