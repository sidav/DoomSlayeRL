package main

type cell struct {
	IsPassable bool
	Appearance rune
}

type dungeon struct {
	player pawn
	tiles  [levelsizex][levelsizey]cell
	pawns  []pawn
}

func (dung *dungeon) initialize_level() { //crap of course
	dung.player = p_createPawn("player", 1, 1)
	dung.pawns = make([]pawn, 0)
	dung.pawns = append(dung.pawns, p_createPawn("zombie", 5, 5))
	dung.pawns = append(dung.pawns, p_createPawn("imp", 3, 1))
	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			dung.tiles[x][y].Appearance = ' '
			dung.tiles[x][y].IsPassable = true
			if x*y == 0 || x == levelsizex-1 || y == levelsizey-1 {
				dung.tiles[x][y].Appearance = '#'
				dung.tiles[x][y].IsPassable = false
			}
		}
	}
}

func (dung *dungeon) isPawnPresent(ix, iy int) bool {
	x, y := dung.player.x, dung.player.y
	if ix == x && iy == y {
		return true
	}
	for i := 0; i < len(dung.pawns); i++ {
		x, y = dung.pawns[i].x, dung.pawns[i].y
		if ix == x && iy == y {
			return true
		}
	}
	return false
}

func (dung *dungeon) getPawnAt(x, y int) *pawn {
	for i := 0; i < len(dung.pawns); i++ {
		px, py := dung.pawns[i].x, dung.pawns[i].y
		if px == x && py == y {
			return &dung.pawns[i]
		}
	}
	return nil
}

func (dung *dungeon) isTilePassable(x, y int) bool {
	return dung.tiles[x][y].IsPassable
}

func (dung *dungeon) isTilePassableAndNotOccupied(x, y int) bool {
	return dung.isTilePassable(x, y) && !dung.isPawnPresent(x, y)
}
