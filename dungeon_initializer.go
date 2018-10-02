package main

import (
	"GoConsoleWrapper/console_wrapper"
	"GoRoguelike/BSP_dungeon_generator"
	"GoRoguelike/routines"
)

func (dung *dungeon) initialize_level() { //crap of course
	dung.player = p_createPlayer(1, 1)
	dung.pawns = make([]*p_pawn, 0)

	dung.MakeMapFromGenerated()

	dung.spawnPawnAtRandomPosition("zombie")
	dung.spawnPawnAtRandomPosition("zombie")
	dung.spawnPawnAtRandomPosition("zombie")
	dung.spawnPawnAtRandomPosition("zombie")
	dung.spawnPawnAtRandomPosition("zombie")
	dung.spawnPawnAtRandomPosition("zombie")
	dung.spawnPawnAtRandomPosition("zombie")
	dung.spawnPawnAtRandomPosition("imp")
	dung.spawnPawnAtRandomPosition("archvile")
	// dung.spawnItemAtRandomPosition("pistol")
	dung.spawnItemAtRandomPosition("clip")
	dung.spawnItemAtRandomPosition("clip")
}

func (dung *dungeon) MakeMapFromGenerated(){
	BSP_dungeon_generator.SetGeneratorRandomSeed(routines.Random(0))
	generated_map := BSP_dungeon_generator.GenerateDungeon(levelsizex, levelsizey, 0, 0, 0, 0, 3)
	for x := 0; x < levelsizex; x++ {
		for y := 0; y < levelsizey; y++ {
			currDungCell := &dung.tiles[x][y]
			currGenCell := generated_map.GetCell(x, y)
			switch currGenCell {
			case '+':
				currDungCell.cCell = &consoleCell{appearance: currGenCell, color: console_wrapper.CYAN}
				currDungCell.IsPassable = true
				currDungCell.opaque = true
			case '~':
				currDungCell.cCell = &consoleCell{appearance: currGenCell, color: console_wrapper.GREEN}
				currDungCell.IsPassable = false
				currDungCell.opaque = false
			case '#':
				currDungCell.cCell = &consoleCell{appearance: currGenCell, color: console_wrapper.BEIGE}
				currDungCell.IsPassable = false
				currDungCell.opaque = true
			default:
				currDungCell.cCell = &consoleCell{appearance: currGenCell, color: console_wrapper.BEIGE}
				currDungCell.IsPassable = true
				currDungCell.opaque = false
			}
		}
	}
}
