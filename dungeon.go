package main

import (
	"GoRoguelike/routines"
)

type dungeon struct {
	player *p_pawn
	tiles  [levelsizex][levelsizey]d_tile
	pawns  []*p_pawn
	items  []*i_item
}

func (dung *dungeon) initialize_level() { //crap of course
	dung.player = p_createPawn("player", 1, 1)
	dung.pawns = make([]*p_pawn, 0)
	dung.items = append(dung.items, i_createItem("clip", 7, 8))
	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			dung.tiles[x][y].Appearance = ' '
			dung.tiles[x][y].IsPassable = true
			if x*y == 0 || x == levelsizex-1 || y == levelsizey-1 || (y == 5 && x != levelsizey-2) {
				dung.tiles[x][y].Appearance = '#'
				dung.tiles[x][y].IsPassable = false
				dung.tiles[x][y].opaque = true
			}
		}
	}
	//dung.pawns = append(dung.pawns, p_createPawn("imp", 1, 9))
	dung.pawns = append(dung.pawns, p_createPawn("imp", 5, 6))
	//dung.spawnPawnAtRandomPosition("zombie")
	//dung.spawnPawnAtRandomPosition("imp")
	// dung.spawnPawnAtRandomPosition("archvile")
	dung.items = append(dung.items, i_createWeapon("pistol", 5, 6))
}

func (dung *dungeon) visibleLineExists(fx, fy, tx, ty int) bool {
	line := routines.GetLine(fx, fy, tx, ty)
	for i := 0; i < len(line); i++ {
		if dung.tiles[line[i].X][line[i].Y].opaque {
			return false
		}
	}
	return true
}

func (d *dungeon) getListOfPawnsVisibleFrom(fx, fy int) []*p_pawn {
	list := make([]*p_pawn, 0)
	for _, pawn := range d.pawns {
		x, y := pawn.x, pawn.y
		if d.visibleLineExists(fx, fy, x, y) {
			list = append(list, pawn)
		}
	}
	return list
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

func (dung *dungeon) getPawnAt(x, y int) *p_pawn {
	px, py := dung.player.x, dung.player.y
	if px == x && py == y {
		return dung.player
	}
	for i := 0; i < len(dung.pawns); i++ {
		px, py = dung.pawns[i].x, dung.pawns[i].y
		if px == x && py == y {
			return dung.pawns[i]
		}
	}
	return nil
}

func (d *dungeon) removePawn(p *p_pawn) {
	for i := 0; i < len(d.pawns); i++ {
		if p == d.pawns[i] {
			d.pawns = append(d.pawns[:i], d.pawns[i+1:]...) // ow it's fucking... magic!
		}
	}
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

func (d *dungeon) addItemToFloor(item *i_item) {
	d.items = append(d.items, item)
}

func (d *dungeon) getListOfItemsAt(ix, iy int) []*i_item {
	items := make([]*i_item, 0)
	for i := 0; i < len(d.items); i++ {
		x, y := d.items[i].x, d.items[i].y
		if ix == x && iy == y {
			items = append(items, d.items[i])
		}
	}
	return items
}

func (d *dungeon) removeItemFromFloor(item *i_item) {
	for i := 0; i < len(d.items); i++ {
		if item == d.items[i] {
			d.items = append(d.items[:i], d.items[i+1:]...) // ow it's fucking... magic!
		}
	}
}

func (dung *dungeon) isTilePassable(x, y int) bool {
	return dung.tiles[x][y].IsPassable
}

func (dung *dungeon) isTilePassableAndNotOccupied(x, y int) bool {
	return dung.isTilePassable(x, y) && !dung.isPawnPresent(x, y)
}
