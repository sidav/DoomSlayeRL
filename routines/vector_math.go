package routines

import "math"

type Vector struct {
	X, Y float64
}

func (v *Vector) InitByStartAndEndInt(sx, sy, ex, ey int) {
	v.X = float64(ex-sx)
	v.Y = float64(ey-sy)
}

func (v *Vector) InitByIntegers(x, y int){
	v.X = float64(x)
	v.Y = float64(y)
}

func (v *Vector) Add(w *Vector){
	v.X += w.X
	v.Y += w.Y
}

func (v *Vector) GetRoundedCoords() (int, int) {
	return int(math.Round(v.X)), int(math.Round(v.Y))
}

func (v *Vector) GetUnitVector() *Vector {
	x, y := v.X, v.Y
	length := math.Sqrt(float64(x*x + y*y))
	newx, newy := x/length, y/length
	return &Vector{newx, newy}
}

func (v *Vector) TransformIntoUnitVector() {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y)
	v.X /= length
	v.Y /= length
}
