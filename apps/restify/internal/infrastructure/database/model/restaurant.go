package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Restaurant struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserID      uuid.UUID `gorm:"index"`
	Name        string
	Address     string
	CoordinateX float64
	CoordinateY float64
	OpenTime    string         `gorm:"time"`
	CloseTime   string         `gorm:"time"`
	WorkDays    pq.StringArray `gorm:"type:text[]"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

func (r Restaurant) OpenTimeHour() uint8 {
	openTimeHour, _ := strconv.Atoi(strings.Split(r.OpenTime, ":")[0])
	return uint8(openTimeHour)
}

func (r Restaurant) OpenTimeMinute() uint8 {
	openTimeMinute, _ := strconv.Atoi(strings.Split(r.OpenTime, ":")[1])
	return uint8(openTimeMinute)
}

func (r Restaurant) CloseTimeHour() uint8 {
	closeTimeHour, _ := strconv.Atoi(strings.Split(r.CloseTime, ":")[0])
	return uint8(closeTimeHour)
}

func (r Restaurant) CloseTimeMinute() uint8 {
	closeTimeMinute, _ := strconv.Atoi(strings.Split(r.CloseTime, ":")[1])
	return uint8(closeTimeMinute)
}
