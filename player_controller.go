package main

func playerControl(d *dungeon) {
	key_pressed := readKey()
	movex := 0
	movey := 0
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
	}

	// log.appendMessage(key_pressed)

	if movex != 0 || movey != 0 {
		moveOrMeleeAttackPawn(&d.player, d, movex, movey)
	}
}
