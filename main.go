package main

import (
	cw "github.com/sidav/golibrl/console"
)

func main() {
	cw.Init_console("DoomSlayeRL", cw.SDLRenderer)
	defer cw.Close_console()
	g := game{}
	g.runGame()
}
