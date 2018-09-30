package main

func (d *dungeon) GetFieldOfVisionFrom(fx, fy int) [levelsizex][levelsizey]bool {
	var final, first, second [levelsizex][levelsizey]bool

	// first stage of the algorithm
	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			if d.visibleLineExists(fx, fy, x, y) {
				first[x][y] = true
			}
		}
	}
	// second stage of the algorithm
	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			if first[x][y] == false {
				if x > 0 && first[x-1][y] {
					second[x][y] = true
				}
				if x < levelsizex-1 && first[x+1][y] {
					second[x][y] = true
				}
				if y > 0 && first[x][y-1] {
					second[x][y] = true
				}
				if y < levelsizey-1 && first[x][y+1] {
					second[x][y] = true
				}
			}
		}
	}

	// merge stages
	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			final[x][y] = first[x][y] || second[x][y]
		}
	}

	return final
}
