package main

import (
	cw "TCellConsoleWrapper/tcell_wrapper"
)

func main() {
	cw.Init_console()
	defer cw.Close_console()
	g := game{}
	g.runGame()
}
