package main

import "fmt"

func plr_playerControl(d *dungeon) {
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
		case "g":
			plr_pickUpItem(d)
		case "f":
			plr_fire(d)
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
		m_moveOrMeleeAttackPawn(d.player, d, movex, movey)
	}
	plr_checkItemsOnFloor(d)
}

func plr_fire(d *dungeon) {
	p := d.player
	if p.weaponInHands == nil {
		log.appendMessage("You have nothing to fire with!")
	}
}

func plr_pickUpItem(d *dungeon) {
	p := d.player
	items := d.getListOfItemsAt(p.x, p.y)
	for i := 0; i < len(items); i++ {
		item := items[i]
		switch items[i].getType() {
		case "weapon":
			p.weaponInHands = item
			d.removeItemFromFloor(items[i])
			log.appendMessage(fmt.Sprintf("You pick up and equip the %s.", p.weaponInHands.name))
			return
		}
	}
	if len(items) == 0 {
		log.appendMessage("There is nothing here.")
		return
	}
	log.appendMessage("Hmm... Can't pick that up.")
}

func plr_checkItemsOnFloor(d *dungeon) {
	px, py := d.player.getCoords()
	items := d.getListOfItemsAt(px, py)
	if len(items) == 1 {
		log.appendMessage(fmt.Sprintf("You see here a %s", items[0].name))
	} else if len(items) > 1 {
		log.appendMessage(fmt.Sprintf("You see here a %s and %d more items", items[0].name, len(items)-1))
	}
}
