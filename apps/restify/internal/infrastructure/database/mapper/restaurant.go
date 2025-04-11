package mapper

import (
	"javifood-restify/internal/domain/entity"
	"javifood-restify/internal/infrastructure/database/model"
)

type RestaurantMapper struct{}

func (rm RestaurantMapper) ToDomain(r model.Restaurant) (*entity.Restaurant, error) {
	workDays := make([]string, len(r.WorkDays))
	for i, day := range r.WorkDays {
		workDays[i] = string(day)
	}
	domainRestaurant, err := entity.NewRestaurant(
		r.ID.String(),
		r.UserID,
		r.Name,
		r.Address,
		r.OpenTime,
		r.CloseTime,
		r.CoordinateX,
		r.CoordinateY,
		workDays,
		r.CreatedAt,
		r.UpdatedAt,
	)
	return domainRestaurant, err
}

func (rm RestaurantMapper) ToDatabase(r entity.Restaurant) *model.Restaurant {
	return &model.Restaurant{
		ID:          r.ID.Value,
		UserID:      r.UserID,
		Name:        r.Name,
		Address:     r.Address,
		CoordinateX: r.Coordinate.X,
		CoordinateY: r.Coordinate.Y,
		OpenTime:    r.OpenTime.String(),
		CloseTime:   r.CloseTime.String(),
		WorkDays:    r.WorkDays.Value,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}
