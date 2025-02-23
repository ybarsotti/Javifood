package valueobject

import (
	"javifood-restify/internal/domain"
	"slices"
)

var AVAILABLE_DAYS = []string{"Monday", "Tuesday", "Wednesday", "Thirsday", "Friday", "Saturday", "Sunday"}

type WorkDays struct {
	Value []string `faker:"slice_len=5"`
}

func NewWorkDays(days []string) (*WorkDays, error) {
	for _, day := range days {
		if !slices.Contains(AVAILABLE_DAYS, day) {
			return nil, domain.InvalidValueError
		}
	} 
	return &WorkDays{
		Value: days,
	}, nil
}
