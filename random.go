package main

const (
	_RANDOM_a = 513
	_RANDOM_c = 313
	_RANDOM_m = 65536
)

var (
	_RANDOM_x int
)

func randomize() {
	_RANDOM_x = 0 // wow, so random
}

func random(modulo int) int {
	_RANDOM_x = (_RANDOM_x*_RANDOM_a + _RANDOM_c) % _RANDOM_m
	return _RANDOM_x % modulo
}

func rollDice(dnum, dval, dmod int) int {
	var result int
	for i := 0; i < dnum; i++ {
		result += random(dval) + 1
	}
	return result + dmod
}
