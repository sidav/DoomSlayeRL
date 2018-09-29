package main

type cell struct {
	IsPassable bool
	Appearance rune
}

type dungeon struct {
	player pawn
	tiles [levelsizex][levelsizey] cell
	pawns []pawn
}

func (dung *dungeon) initialize_level() {
	dung.player = pawn{appearance:'@', x:1, y:1}
	dung.pawns = make([]pawn, 0)
	dung.pawns = append(dung.pawns, pawn{'@', 1, 1, 1})
	for x:=0; x < levelsizex; x++ {
		for y:=0; y<levelsizey; y++{
			dung.tiles[x][y].Appearance = ' '
			if x * y == 0 || x == levelsizex-1 || y == levelsizey - 1 {
				dung.tiles[x][y].Appearance = '#'
			}
		}
	}
}