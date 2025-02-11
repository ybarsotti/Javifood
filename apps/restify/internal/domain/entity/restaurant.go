package entity

import "time"

type Restaurant struct {
	ID         string
	UserID     string
	Name       string
	Address    string
	Coordinate *Coordinate
	OpenTime   time.Time
	Closetime  time.Time
	WorkDays   *WorkDays
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewRestaurant(id, userId, name, address, coordX, coordY string, openTime, closeTime time.Time, workdays []string) (*Restaurant, error) {
	coordinate, err := NewCoordinate(coordX, coordX) 
	if err != nil {
		return nil, err
	}
	workDays, err := NewWorkDays(workdays)
	if err != nil {
		return nil, err
	}

	return &Restaurant{
		ID: id,
		UserID: userId,
		Name: name,
		Address: address,
		Coordinate: coordinate,
		OpenTime: openTime,
		Closetime: closeTime,
		WorkDays: workDays,
		CreatedAt: time.Now(),
	}, nil
}
