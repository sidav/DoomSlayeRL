package main

import "DoomSlayeRL/routines"

type dice struct {
	dnum, dval, dmod int
}

func (d *dice) roll() int {
	return routines.RollDice(d.dnum, d.dval, d.dmod)
}
