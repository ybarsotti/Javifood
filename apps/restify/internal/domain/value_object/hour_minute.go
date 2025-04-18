package valueobject

import (
	"fmt"
	"javifood-restify/internal/domain"
	"javifood-restify/pkg/utils"
)

type HourMinute struct {
	Hour   int
	Minute int
}

// Time format is 24h
func NewHourMinute(t string) (*HourMinute, error) {
	hour, minute := utils.SplitStringTimeIntoHourMinute(t)
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
