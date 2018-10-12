package BSP_dungeon_generator

const (
	r_a = 513
	r_c = 313
	r_m = 262147
)

var (
	r_x int
)

func SetGeneratorRandomSeed(seed int) {
	r_x = seed
}

func random(modulo int) int {
	r_x = (r_x*r_a + r_c) % r_m
	return r_x % modulo
}


func randInRange(from, to int) int { //should be inclusive
	if to < from {
		t := from
		from = to
		to = t
	}
	if from == to {
		return from
	}
	return random(to-from+1) + from // TODO: replace routines.random usage with package own implementation
}
