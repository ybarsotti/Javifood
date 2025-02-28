package model

import (
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
