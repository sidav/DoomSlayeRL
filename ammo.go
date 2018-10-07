package main

type AMMO_TYPE byte

const (
	AMMO_BULL AMMO_TYPE = 0
	AMMO_SHEL AMMO_TYPE = 1
	AMMO_RCKT AMMO_TYPE = 2
	AMMO_CELL AMMO_TYPE = 3
)

type i_ammoData struct {
	ammo [4]int
}
