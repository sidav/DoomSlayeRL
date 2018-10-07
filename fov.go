package main

import "GoRoguelike/routines"

const DEFAULT_SIGHT_RANGE = 9

func (dung *dungeon) unobstructedLineExists(fx, fy, tx, ty int) bool { // visible AND free of pawns line
	line := routines.GetLine(fx, fy, tx, ty)
	for i := 1; i < len(line)-1; i++ { // we skip first and last cells
		if dung.isTileOpaque(line[i].X, line[i].Y) || dung.isPawnPresent(line[i].X, line[i].Y) {
			return false
		}
	}
	return true
}

func (d *dungeon) canPawnSeeCoords(p *p_pawn, tx, ty int) bool {
	fx, fy := p.getCoords()
	sRange := p.sightRange
	if sRange == 0 {
		sRange = DEFAULT_SIGHT_RANGE // Need a better solution for this...
	}
	if !areCoordinatesInRangeFrom(fx, fy, tx, ty, sRange) {
		return false
	}
	line := routines.GetLine(fx, fy, tx, ty)
	for i := 1; i < len(line); i++ { // we skip first cell
		if d.isTileOpaque(line[i].X, line[i].Y) {
			return false
		}
	}
	return true
}

func (d *dungeon) getListOfPawnsVisibleFor(seer *p_pawn) []*p_pawn {
	list := make([]*p_pawn, 0)
	for _, pawn := range d.pawns {
		x, y := pawn.x, pawn.y
		if d.canPawnSeeCoords(seer, x, y) {
			list = append(list, pawn)
		}
	}
	return list
}

func (d *dungeon) GetFieldOfVisionFor(seer *p_pawn) [levelsizex][levelsizey]bool {
	var final, first, second [levelsizex][levelsizey]bool

	sRange := seer.sightRange
	if sRange == 0 {
		sRange = DEFAULT_SIGHT_RANGE // Need a better solution for this...
	}
	seerx, seery := seer.getCoords()

	// first stage of the algorithm
	for x := seerx - sRange; x <= seerx+sRange; x++ {
		for y := seery - sRange; y <= seery+sRange; y++ {
			if !areCoordinatesValid(x, y) || !areCoordinatesInRangeFrom(seerx, seery, x, y, sRange) {
				continue
			}
			if d.canPawnSeeCoords(seer, x, y) {
				first[x][y] = true
			}
		}
	}
	// second stage of the algorithm
	for x := seerx - sRange; x <= seerx+sRange; x++ {
		for y := seery - sRange; y <= seery+sRange; y++ {
			if !areCoordinatesValid(x, y) || !areCoordinatesInRangeFrom(seerx, seery, x, y, sRange) {
				continue
			}
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
