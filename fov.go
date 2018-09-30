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
			neighbourCheck:
				for i := -1; i < 2; i++ {
					for j := -1; j < 2; j++ {
						if x+i >= 0 && x+i < levelsizex && y+j >= 0 && y+j < levelsizey {
							if first[x+i][y+j] {
								second[x][y] = true
								break neighbourCheck
							}
						}
					}
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
