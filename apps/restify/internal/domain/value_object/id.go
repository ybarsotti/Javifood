package valueobject

import "github.com/google/uuid"

type ID struct {
	Value uuid.UUID
}

func NewID(value string) (*ID, error) {
	if value != "" {
		parsedUUID, err := uuid.Parse(value)
		if err != nil {
			return nil, err
		}
		return &ID{
			Value: parsedUUID,
		}, nil
	}

	newUUID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &ID{
		Value: newUUID,
	}, nil
}
