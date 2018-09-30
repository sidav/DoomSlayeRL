package routines

import (
	"time"
)

const (
	a = 513
	c = 313
	m = 65536
)

var (
	x int
)

func Randomize() {
	x = int(time.Duration(time.Now().UnixNano())/time.Millisecond) % m
}

func Random(modulo int) int {
	x = (x*a + c) % m
	return x % modulo
}

func RollDice(dnum, dval, dmod int) int {
	var result int
	for i := 0; i < dnum; i++ {
		result += Random(dval) + 1
	}
	return result + dmod
}
