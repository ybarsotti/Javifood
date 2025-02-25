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
		r.UserID.String(),
		r.Name,
		r.Address,
		r.CoordinateX,
		r.CoordinateY,
		r.OpenTimeHour(),
		r.OpenTimeMinute(),
		r.CloseTimeHour(),
		r.CloseTimeMinute(),
		workDays,
		r.CreatedAt,
		r.UpdatedAt,
	)
	return domainRestaurant, err
}

func (rm RestaurantMapper) ToDatabase(r entity.Restaurant) (*model.Restaurant, error) {
	return nil, nil
}
