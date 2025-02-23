package entity_test

import (
	"javifood-restify/internal/domain/entity"
	"slices"
	"testing"
	"time"
)

func TestRestaurant_NewRestaurantIsValid(t *testing.T) {
	restaurantID := "0194fc3f-7a6d-79fe-ab07-7426c04720cf"
	userID := "0194fc3f-7a6d-79fe-ab07-7426c04720ca"
	name := "Restaurant 1"
	address := "Address 1"
	coordinateX := -65.87425
	coordinateY := 25.23404
	openHour := uint8(8)
	openMinute := uint8(00)
	closeHour := uint8(21)
	closeMinute := uint8(00)
	workDays := []string{"Monday", "Tuesday"}
	restaurant, err := entity.NewRestaurant(restaurantID, userID, name, address, coordinateX, coordinateY,
		openHour, openMinute, closeHour, closeMinute, workDays, time.Now(), time.Now())
	if err != nil {
		t.Errorf("not expected error: %v", err.Error())
	}
	if restaurant.ID.Value.String() != restaurantID {
		t.Errorf("restaurant id should not be empty")
	}
	if restaurant.UserID.Value.String() != userID {
		t.Errorf("user id should not be empty")
	}
	if restaurant.Name != name {
		t.Errorf("name does not match ")
	}
	if restaurant.Address != address {
		t.Errorf("address does not match ")
	}
	if restaurant.Coordinate.X != coordinateX {
		t.Errorf("coordinateX does not match ")
	}
	if restaurant.Coordinate.Y != coordinateY {
		t.Errorf("coordinateY does not match ")
	}
	if restaurant.OpenTime.String() != "08:00" {
		t.Errorf("opentime does not match")
	}
	if restaurant.CloseTime.String() != "21:00" {
		t.Errorf("closetime does not match")
	}
	if !slices.Equal(restaurant.WorkDays.Value, workDays) {
		t.Errorf("workdays does not match")
	}
}

func TestRestaurant_ValidateMissingUserID(t *testing.T) {
}
