package main

import "GoRoguelike/routines"

type cell struct {
	IsPassable bool
	Appearance rune
}

type dungeon struct {
	player p_pawn
	tiles  [levelsizex][levelsizey]cell
	pawns  []p_pawn
	items  []i_item
}

func (dung *dungeon) initialize_level() { //crap of course
	dung.player = p_createPawn("player", 1, 1)
	dung.pawns = make([]p_pawn, 0)
	dung.items = append(dung.items, i_createItem("clip", 7, 8))
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
	dung.spawnPawnAtRandomPosition("zombie")
	dung.spawnPawnAtRandomPosition("imp")
	dung.spawnPawnAtRandomPosition("archvile")
	dung.items = append(dung.items, i_createWeapon("pistol", 4, 5))
}

func (dung *dungeon) spawnPawnAtRandomPosition(name string) {
	for tries := 0; tries < 1000; tries++ {
		x, y := routines.Random(levelsizex), routines.Random(levelsizey)
		if dung.isTilePassableAndNotOccupied(x, y) {
			dung.pawns = append(dung.pawns, p_createPawn(name, x, y))
			return
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

func (d *dungeon) isItemPresent(ix, iy int) bool {
	for i := 0; i < len(d.items); i++ {
		x, y := d.items[i].x, d.items[i].y
		if ix == x && iy == y {
			return true
		}
	}
	return false
}

func (d *dungeon) getItemAt(ix, iy int) *i_item {
	for i := 0; i < len(d.items); i++ {
		x, y := d.items[i].x, d.items[i].y
		if ix == x && iy == y {
			return &d.items[i]
		}
	}
	return nil
}

func (dung *dungeon) getPawnAt(x, y int) *p_pawn {
	px, py := dung.player.x, dung.player.y
	if px == x && py == y {
		return &dung.player
	}
	for i := 0; i < len(dung.pawns); i++ {
		px, py = dung.pawns[i].x, dung.pawns[i].y
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
