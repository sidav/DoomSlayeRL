package main

import (
	// "DoomSlayeRL/BSP_dungeon_generator"
	dungen "CyclicDungeonGenerator/layout_generation"
	dungToTiled "CyclicDungeonGenerator/layout_to_tiled"
	"DoomSlayeRL/routines"
	cw "github.com/sidav/golibrl/console"
)

func (dung *dungeon) initialize_level() { //crap of course
	dung.MakeMapFromGenerated()

	dung.spawnPlayerAtRandomPosition()

	dung.pawns = make([]*p_pawn, 0)

	dung.init_placeItemsAndEnemies()
}

func (dung *dungeon) initTilesArrayForSize() {
	dung.tiles = make([][]d_tile, levelsizex)
	for i := range dung.tiles {
		dung.tiles[i] = make([]d_tile, levelsizey)
	}
}

func (dung *dungeon) init_placeItemsAndEnemies() {
	dung.spawnPawnAtRandomPosition("unwilling", 15)
	dung.spawnPawnAtRandomPosition("zombie soldier", 12)
	dung.spawnPawnAtRandomPosition("zombie sergeant", 12)
	dung.spawnPawnAtRandomPosition("heavy weapon dude", 10)
	dung.spawnPawnAtRandomPosition("imp", 7)
	dung.spawnPawnAtRandomPosition("pinky", 7)
	dung.spawnPawnAtRandomPosition("hellknight", 4)

	dung.spawnItemAtRandomPosition("clip", 15)
	dung.spawnItemAtRandomPosition("cell", 5)
	dung.spawnItemAtRandomPosition("shells", 15)
	dung.spawnItemAtRandomPosition("ammunition crate", 3)

	dung.spawnItemAtRandomPosition("stimpack", 15)
	dung.spawnItemAtRandomPosition("small medikit", 7)
	dung.spawnItemAtRandomPosition("large medikit", 4)
	dung.spawnItemAtRandomPosition("soulsphere", 2)

	dung.spawnItemAtRandomPosition("green armor", 3)
	dung.spawnItemAtRandomPosition("red armor", 2)
	dung.spawnItemAtRandomPosition("blue armor", 1)

	dung.spawnItemAtRandomPosition("chaingun", 1)
	dung.spawnItemAtRandomPosition("heavy pistol", 3)
	dung.spawnItemAtRandomPosition("shotgun", 4)
	dung.spawnItemAtRandomPosition("super shotgun", 2)
	dung.spawnItemAtRandomPosition("assault rifle", 3)
	dung.spawnItemAtRandomPosition("Pancor Jackhammer", 1)
	dung.spawnItemAtRandomPosition("bolt-action rifle", 3)
	dung.spawnItemAtRandomPosition("gauss rifle", 1)
}

func (dung *dungeon) MakeMapFromGenerated() {
	// BSP_dungeon_generator.SetGeneratorRandomSeed(routines.Random(0))
	layout, _ := dungen.Generate(-1, 6, 6) //BSP_dungeon_generator.GenerateDungeon(levelsizex, levelsizey, 7, 60, 0, 50, 5)
	generated_map := dungToTiled.GetTileMap(layout)

	levelsizex = len(*generated_map)
	levelsizey = len((*generated_map)[0])
	dung.initTilesArrayForSize()

	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			currDungCell := &dung.tiles[x][y]
			currGenCell := (*generated_map)[x][y].Char //GetCell(x, y)
			switch currGenCell {
			case '+':
				currDungCell.cCell = &consoleCell{appearance: 16*12+14, color: cw.DARK_CYAN}
				currDungCell.opaque = true
				currDungCell.doorData = &d_doorData{chrForOpened: '\''}
			case '~':
				currDungCell.cCell = &consoleCell{appearance: currGenCell, color: cw.DARK_GREEN}
				currDungCell.IsPassable = false
				currDungCell.opaque = false
			case '#', rune(0):
				currDungCell.cCell = &consoleCell{appearance: 16*11+1, color: cw.BEIGE}
				currDungCell.IsPassable = false
				currDungCell.opaque = true
			default:
				currDungCell.cCell = &consoleCell{appearance: currGenCell, color: cw.BEIGE}
				currDungCell.IsPassable = true
				currDungCell.opaque = false
			}
		}
	}
}


func (dung *dungeon) spawnPlayerAtRandomPosition() {
	for tries := 0; tries < 1000; tries++ {
		x, y := routines.Random(levelsizex), routines.Random(levelsizey)
		if dung.isTilePassable(x, y) {
			dung.player = p_createPlayer(x, y)
			break
		}
	}
}

func (dung *dungeon) spawnPawnAtRandomPosition(name string, count int) { //spawns outside of player's FOV
	for num := 0; num < count; num++ {
		for tries := 0; tries < 1000; tries++ {
			x, y := routines.Random(levelsizex), routines.Random(levelsizey)
			if dung.isTilePassableAndNotOccupied(x, y) && !dung.canPawnSeeCoords(dung.player, x, y) {
				dung.pawns = append(dung.pawns, p_createPawn(name, x, y))
				break
			}
		}
	}
}

func (dung *dungeon) spawnItemAtRandomPosition(name string, count int) {
	for num := 0; num < count; num++ {
		for tries := 0; tries < 1000; tries++ {
			x, y := routines.Random(levelsizex), routines.Random(levelsizey)
			if dung.isTilePassableAndNotOccupied(x, y) {
				dung.items = append(dung.items, i_createItem(name, x, y))
				break
			}
		}
	}
}
