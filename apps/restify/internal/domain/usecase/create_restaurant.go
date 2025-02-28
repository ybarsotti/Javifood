package usecase

import (
	"context"
)

type CreateRestaurantUseCaseInputDto struct {
	UserID      string
	Name        string
	Address     string
	CoordinateX float64
	CoordinateY float64
	OpenTime    string
	CloseTime   string
	WorkDays    []string
}

func NewCreateRestaurantInputDto(
	userId, name, address, openTime, closeTime string,
	coordinateX, coordinateY float64,
	workDays []string,
) *CreateRestaurantUseCaseInputDto {
	return &CreateRestaurantUseCaseInputDto{
		UserID:      userId,
		Name:        name,
		Address:     address,
		CoordinateX: coordinateX,
		CoordinateY: coordinateY,
		OpenTime:    openTime,
		CloseTime:   closeTime,
		WorkDays:    workDays,
	}
}

type CreateRestaurantUseCase interface {
	Execute(ctx context.Context, input CreateRestaurantUseCaseInputDto) error
}
