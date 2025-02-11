package entity

type Coordinate struct {
	X string
	Y string
}

func NewCoordinate(x, y string) (*Coordinate, error) {
	return &Coordinate{
		x, y,
	}, nil
}
