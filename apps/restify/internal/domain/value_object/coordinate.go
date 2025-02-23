package valueobject

type Coordinate struct {
	X float64 `faker:"lat"`
	Y float64 `faker:"long"`
}

func NewCoordinate(x, y float64) (*Coordinate, error) {
	return &Coordinate{
		x, y,
	}, nil
}
