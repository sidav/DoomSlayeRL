package main

import "fmt"

func i_createItem(name string, x, y int) i_item {
	var i i_item
	switch name {
	case "clip":
		i = i_item{appearance: '"', name: name}
	default:
		i = i_item{appearance: '?', name: "UNKNOWN ITEM " + name}
	}
	i.x = x
	i.y = y
	return i
}

func i_createCorpseFor(p *pawn) i_item {
	x, y := p.x, p.y
	name := fmt.Sprintf("a %s corpse", p.name)
	return i_item{name: name, x: x, y: y, appearance: '%'}
}
