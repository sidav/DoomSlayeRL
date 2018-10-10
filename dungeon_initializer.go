package main

import (
	"DoomSlayeRL/BSP_dungeon_generator"
	"DoomSlayeRL/routines"
	cw "TCellConsoleWrapper/tcell_wrapper"
)

func (dung *dungeon) initialize_level() { //crap of course
	dung.player = p_createPlayer(1, 1)
	dung.pawns = make([]*p_pawn, 0)

	dung.MakeMapFromGenerated()
	dung.init_placeItemsAndEnemies()
}

func (dung *dungeon) init_placeItemsAndEnemies() {
	for i := 0; i < 25; i++ {
		// dung.spawnPawnAtRandomPosition("zombie")
		dung.spawnPawnAtRandomPosition("imp")
	}
	for i := 0; i < 5; i++ {
		dung.spawnItemAtRandomPosition("stimpack")
		dung.spawnItemAtRandomPosition("pistol")
		dung.spawnItemAtRandomPosition("clip")
		dung.spawnItemAtRandomPosition("cell")
		dung.spawnItemAtRandomPosition("shells")
		dung.spawnItemAtRandomPosition("ammunition crate")
	}
	dung.spawnItemAtRandomPosition("soulsphere")
	dung.spawnItemAtRandomPosition("chaingun")
	dung.spawnItemAtRandomPosition("shotgun")
	dung.spawnItemAtRandomPosition("super shotgun")
	dung.spawnItemAtRandomPosition("assault rifle")
	dung.spawnItemAtRandomPosition("Pancor Jackhammer")
	dung.spawnItemAtRandomPosition("bolt-action rifle")
	dung.spawnItemAtRandomPosition("gauss rifle")
}

func (dung *dungeon) MakeMapFromGenerated() {
	BSP_dungeon_generator.SetGeneratorRandomSeed(routines.Random(0))
	generated_map := BSP_dungeon_generator.GenerateDungeon(levelsizex, levelsizey, 5, 0, 0, 40, 5)
	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			currDungCell := &dung.tiles[x][y]
			currGenCell := generated_map.GetCell(x, y)
			switch currGenCell {
			case '+':
				currDungCell.cCell = &consoleCell{appearance: '╬', color: cw.DARK_CYAN}
				currDungCell.IsPassable = true
				currDungCell.opaque = true
			case '~':
				currDungCell.cCell = &consoleCell{appearance: currGenCell, color: cw.DARK_GREEN}
				currDungCell.IsPassable = false
				currDungCell.opaque = false
			case '#':
				currDungCell.cCell = &consoleCell{appearance: '▒', color: cw.BEIGE}
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
