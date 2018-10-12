package main

import "fmt"

func plr_pickUpAnItem(item *i_item, d *dungeon) {
	p := d.player
	itemShouldBeRemoved := false
	if item.instantlyPickupable {
		itemShouldBeRemoved = p.tryApplyItem(item)
	} else {
		if len(p.inventory.items) >= p.inventory.maxItems {
			log.appendMessage("Your inventory is full.")
			return
		}
		p.inventory.addItem(item)
		itemShouldBeRemoved = true
		log.appendMessagef("You pick up the %s.", item.name)
	}
	if itemShouldBeRemoved {
		d.removeItemFromFloor(item)
	}
}

func plr_UseItemFromInventory(p *p_pawn) {
	item := p.inventory.selectItem()
	if item == nil {
		return
	}
	if p.tryApplyItem(item) {
		p.inventory.removeItem(item)
	}
}

func (p *p_pawn) tryApplyItem(item *i_item) bool {
	switch item.getType() {
	case "ammo":
		if p.inventory.canAmmoBeAdded(item) {
			p.inventory.addItem(item)
			log.appendMessagef("You pick up the %s.", item.name)
			return true
		}
	case "weapon":
		if p.weaponInHands != nil {
			p.inventory.addItem(p.weaponInHands)
		}
		p.weaponInHands = item
		log.appendMessage(fmt.Sprintf("You equip the %s.", p.weaponInHands.name))
		return true
	case "medical":
		if p.hp < p.maxhp || item.medicalData.ignoresMaximum {
			p.hp += item.medicalData.healAmount
			if !item.medicalData.ignoresMaximum && p.hp > p.maxhp {
				p.hp = p.maxhp
			}
			log.appendMessage(fmt.Sprintf("The %s heals you.", item.name))
			return true
		}
	case "armor":
		p.wearedArmor = item
		p.wearedArmor.armorData.currArmor = item.armorData.maxArmor
		log.appendMessage(fmt.Sprintf("You wear the %s.", item.name))
		return true
	default:
		log.appendMessage("WTF is that item?!")
		return false
	}
	return false
}
