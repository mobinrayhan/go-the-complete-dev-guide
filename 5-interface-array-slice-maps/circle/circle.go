package circle

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pow(math.Pi*c.Radius, 2)
}
