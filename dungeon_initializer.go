package main

import (
	"GoConsoleWrapper/console_wrapper"
	"GoRoguelike/BSP_dungeon_generator"
	"GoRoguelike/routines"
)

func (dung *dungeon) initialize_level() { //crap of course
	dung.player = p_createPawn("player", 1, 1)
	dung.pawns = make([]*p_pawn, 0)
	dung.items = append(dung.items, i_createItem("clip", 7, 8))
	dung.MakeMapFromGenerated()
	//dung.pawns = append(dung.pawns, p_createPawn("imp", 1, 9))
	dung.pawns = append(dung.pawns, p_createPawn("imp", 8, 8))
	dung.spawnPawnAtRandomPosition("zombie")
	dung.spawnPawnAtRandomPosition("imp")
	// dung.spawnPawnAtRandomPosition("archvile")
	dung.items = append(dung.items, i_createWeapon("pistol", 1, 2))
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
