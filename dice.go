package main

import "GoRoguelike/routines"

type dice struct {
	dnum, dval, dmod int
}

func (d *dice) roll() int {
	return routines.RollDice(d.dnum, d.dval, d.dmod)
}
