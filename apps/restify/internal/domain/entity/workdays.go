package entity

import "slices"

var AVAILABLE_DAYS = [7]string{"Monday", "Tuesday", "Wednesday", "Thirsday", "Friday", "Saturday", "Sunday"}

type WorkDays struct {
	days []string
}

func NewWorkDays(days []string) (*WorkDays, error) {
	for day := range days {
		if !slices.Contains(AVAILABLE_DAYS, day) {
			return nil, InvalidValueError
		}
	} 
	return &WorkDays{
		days: days,
	}, nil
}
