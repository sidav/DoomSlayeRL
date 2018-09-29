package main

import "GoConsoleWrapper/console_wrapper"

func playerControl(d *dungeon) {
	key_pressed := console_wrapper.Read_key_char()
	movex := 0
	movey := 0
	switch key_pressed {
	case 's':
		movey = 1
	case 'w':
		movey = -1
	case 'a':
		movex = -1
	case 'd':
		movex = 1
	}
	if movex != 0 || movey != 0 {
		movePawn(&d.player, d, movex, movey)
	}
}
