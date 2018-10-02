package main

import (
	"fmt"
)

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
		case "[":	// debug
			RENDER_DISABLE_LOS = !RENDER_DISABLE_LOS
			log.appendMessage("Changed LOS setting.")
		default:
			valid_key_pressed = false
			log.appendMessagef("Unknown key %s (Wrong keyboard layout?)", key_pressed)
			renderLevel(d, true)
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
		return
	}
	targets := d.getListOfPawnsVisibleFrom(p.x, p.y)
	curr_target_index := 0
	// choose target
	if len(targets) > 0 {
		log.appendMessagef("You target with your %s.", p.weaponInHands.name)
		aimx, aimy := targets[curr_target_index].x, targets[curr_target_index].y
	aimLoop:
		for {
			renderLevel(d, false)
			renderLine('*', p.x, p.y, aimx, aimy, true, true)
			keypressed := readKey()
			switch keypressed {
			case "n":
				curr_target_index++
				if curr_target_index >= len(targets) {
					curr_target_index = 0
				}
				if len(targets) > 0 {
					aimx, aimy = targets[curr_target_index].x, targets[curr_target_index].y
				}
			case "w":
				aimy -= 1
			case "s":
				aimy += 1
			case "a":
				aimx -= 1
			case "d":
				aimx += 1
			case "f":
				log.appendMessage("Haha, firing is not implemented yet")
				break aimLoop
			case "ESCAPE":
				log.appendMessage("Okay, then.")
				break aimLoop
			}
		}
	} else {
		log.appendMessage("No targets in sight.")
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
		case "ammo":
			p.inventory.addItem(item)
			log.appendMessage(fmt.Sprintf("You pick up the %s.", item.name))
			d.removeItemFromFloor(item)
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
