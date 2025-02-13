package domain

import (
	"errors"
	"fmt"
)

var (
	InvalidValueError = errors.New("invalid value")
	UserAlreadyHaveRestaurantError = errors.New("current user already have a restaurant")
	InternalServerError = errors.New("internal server error")
	RequiredFieldMissingError = errors.New("required field missing")
	HourMinuteRangeError = errors.New("hour time range is invalid")
)

func NewRequiredFieldMissingError (field string) error {
	return fmt.Errorf("%w: %s", RequiredFieldMissingError, field)
}
