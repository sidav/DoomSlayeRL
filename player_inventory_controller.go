package main

import "fmt"

func plr_pickUpAnItem(item *i_item, d *dungeon){
	p := d.player
	switch item.getType() {
	case "weapon":
		if len(p.inventory.items) >= p.inventory.maxItems {
			log.appendMessage("Your inventory is full.")
			return
		}
		if p.weaponInHands != nil {
			p.inventory.addItem(p.weaponInHands)
		}
		p.weaponInHands = item
		d.removeItemFromFloor(item)
		log.appendMessage(fmt.Sprintf("You pick up and equip the %s.", p.weaponInHands.name))
		return
	case "ammo":
		if p.inventory.canAmmoBeAdded(item) {
			p.inventory.addItem(item)
			log.appendMessage(fmt.Sprintf("You pick up the %s.", item.name))
			d.removeItemFromFloor(item)
			return
		}
	case "medical":
		if item.instantlyPickupable {
			if p.hp < p.maxhp || item.medicalData.ignoresMaximum {
				p.hp += item.medicalData.healAmount
				if !item.medicalData.ignoresMaximum && p.hp > p.maxhp {
					p.hp = p.maxhp
				}
				d.removeItemFromFloor(item)
				log.appendMessage(fmt.Sprintf("The %s heals you.", item.name))
			}
		} else {
			if len(p.inventory.items) >= p.inventory.maxItems {
				log.appendMessage("Your inventory is full.")
				return
			}
			log.appendMessage(fmt.Sprintf("You pick up the %s.", item.name))
			p.inventory.addItem(item)
			d.removeItemFromFloor(item)
		}
		return
	default:
		log.appendMessage("Hmm... Can't pick that up.")
	}
}

func plr_UseItemFromInventory(p *p_pawn){
	item := p.inventory.selectItem()
	if item == nil {
		return
	}
	switch item.getType() {
	case "weapon":
		if p.weaponInHands != nil {
			p.inventory.addItem(p.weaponInHands)
		}
		p.weaponInHands = item
		log.appendMessage(fmt.Sprintf("You equip the %s.", p.weaponInHands.name))
	case "medical":
		if p.hp < p.maxhp || item.medicalData.ignoresMaximum {
			p.hp += item.medicalData.healAmount
			if !item.medicalData.ignoresMaximum && p.hp > p.maxhp {
				p.hp = p.maxhp
			}
			log.appendMessage(fmt.Sprintf("The %s heals you.", item.name))
		}
	default:
		log.appendMessage("Hmm... Can't pick that up.")
	}
	p.inventory.removeItem(item)
}

//func (p *p_pawn) plr_applyItem(item *i_item) {
//	switch item.getType() {
//	case "ammo":
//		p.inventory.addItem(item)
//		log.appendMessage(fmt.Sprintf("You pick up the %s.", item.name))
//		return
//	case "medical":
//		if p.hp < p.maxhp || item.medicalData.ignoresMaximum {
//			p.hp += item.medicalData.healAmount
//			if !item.medicalData.ignoresMaximum && p.hp > p.maxhp {
//				p.hp = p.maxhp
//			}
//			log.appendMessage(fmt.Sprintf("The %s heals you.", item.name))
//		}
//		return
//	default:
//		log.appendMessage("WTF is that item?!")
//	}
//}
