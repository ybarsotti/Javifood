package valueobject

import (
	"fmt"
	"javifood-restify/internal/domain"
)

type HourMinute struct {
	Hour   uint8
	Minute uint8
}

// Time format is 24h
func NewHourMinute(hour, minute uint8) (*HourMinute, error) {
	if hour < 0 || hour > 24 || minute < 0 || minute > 60 {
		return nil, domain.HourMinuteRangeError
	}
	return &HourMinute{
		Hour:   hour,
		Minute: minute,
	}, nil
}

func (hm *HourMinute) String() string {
	return fmt.Sprintf("%02d:%02d", hm.Hour, hm.Minute)
}
