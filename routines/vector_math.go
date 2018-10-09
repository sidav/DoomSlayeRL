package routines

import "math"

type Vector struct {
	X, Y float64
}

func (v *Vector) GetUnitVector() *Vector {
	x, y := v.X, v.Y
	length := math.Sqrt(float64(x*x + y*y))
	newx, newy := x/length, y/length
	return &Vector{newx, newy}
}
