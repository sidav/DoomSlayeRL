package main

import "fmt"

func playerControl(d *dungeon) {
	valid_key_pressed := false
	movex := 0
	movey := 0
	for !valid_key_pressed {
		key_pressed := readKey()
		valid_key_pressed = true
		switch key_pressed {
		case "s":
			movey = 1
		case "w":
			movey = -1
		case "a":
			movex = -1
		case "d":
			movex = 1
		case "ESCAPE":
			GAME_IS_RUNNING = false
		default:
			valid_key_pressed = false
			log.appendMessage(fmt.Sprintf("Unknown key %s (Wrong keyboard layout?)", key_pressed))
			renderLevel(d)
		}
	}
	// log.appendMessage(key_pressed)

	if movex != 0 || movey != 0 {
		m_moveOrMeleeAttackPawn(&d.player, d, movex, movey)
	}
}
