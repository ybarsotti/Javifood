package entity

import (
	"fmt"
	"javifood-restify/internal/domain"
	valueobject "javifood-restify/internal/domain/value_object"
	"time"

	"github.com/google/uuid"
)

type Restaurant struct {
	ID         *valueobject.ID         `faker:"uuid_digit"`
	UserID     *valueobject.ID         `faker:"uuid_digit"`
	Name       string                  `faker:"name"`
	Address    string                  `faker:"real_address"`
	Coordinate *valueobject.Coordinate `faker:"coordinate"`
	OpenTime   *valueobject.HourMinute `faker:"hourminute"`
	CloseTime  *valueobject.HourMinute `faker:"hourminute"`
	WorkDays   *valueobject.WorkDays   `faker:""`
	CreatedAt  *time.Time              `faker:"date"`
	UpdatedAt  *time.Time              `faker:"date"`
}

func NewRestaurant(
	id, userId, name, address string,
	coordX, coordY float64,
	openTimeHour, openTimeMinute, closeTimeHour, closeTimeMinute uint8,
	workdays []string,
	createdAt, updatedAt *time.Time,
) (*Restaurant, error) {
	coordinate, err := valueobject.NewCoordinate(coordX, coordY)
	if err != nil {
		return nil, err
	}
	workDays, err := valueobject.NewWorkDays(workdays)
	if err != nil {
		return nil, err
	}
	rid, err := valueobject.NewID(id)
	if err != nil {
		return nil, err
	}
	uid, err := valueobject.NewID(userId)
	if err != nil {
		return nil, err
	}
	openTime, err := valueobject.NewHourMinute(openTimeHour, openTimeMinute)
	if err != nil {
		return nil, err
	}
	closeTime, err := valueobject.NewHourMinute(closeTimeHour, closeTimeMinute)
	if err != nil {
		return nil, err
	}
	restaurant := &Restaurant{
		ID:         rid,
		UserID:     uid,
		Name:       name,
		Address:    address,
		Coordinate: coordinate,
		OpenTime:   openTime,
		CloseTime:  closeTime,
		WorkDays:   workDays,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}
	if err = restaurant.Validate(); err != nil {
		return nil, fmt.Errorf("error while validating restaurant: %v", err)
	}
	return restaurant, nil
}

func (r *Restaurant) Validate() error {
	if r.UserID.Value == uuid.Nil {
		return domain.NewRequiredFieldMissingError("userID")
	}
	if r.Name == "" {
		return domain.NewRequiredFieldMissingError("name")
	}
	if r.Address == "" {
		return domain.NewRequiredFieldMissingError("address")
	}
	if r.Coordinate == nil {
		return domain.NewRequiredFieldMissingError("coordinate")
	}
	if r.OpenTime == nil {
		return domain.NewRequiredFieldMissingError("opentime")
	}
	if r.CloseTime == nil {
		return domain.NewRequiredFieldMissingError("closetime")
	}
	if len(r.WorkDays.Value) == 0 {
		return domain.NewRequiredFieldMissingError("workdays")
	}
	return nil
}
