package valueobject_test

import (
	valueobject "javifood-restify/internal/domain/value_object"
	"slices"
	"testing"
)

func TestWorkDays_IsValid(t *testing.T) {
	days := []string{"Monday", "Friday"}
	workDays, err := valueobject.NewWorkDays(days)
	for _, day := range days {
		if !slices.Contains(workDays.Value, day) {
			t.Errorf("should have given day %s", day)
		}
	}
	if err != nil {
		t.Errorf("should not raise error")
	}
}

func TestWorkDays_IsInvalid(t *testing.T) {
	_, err := valueobject.NewWorkDays([]string{"Invalid", "Friday"})
	if err == nil {
		t.Errorf("should raise error")
	}
}
