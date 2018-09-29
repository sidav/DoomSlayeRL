package main

import (
	cw "GoConsoleWrapper/console_wrapper"
)

func main() {
	cw.Init_console()
	defer cw.Close_console()
	go func() {
		cw.Run_event_listener()
	}()

	g := game{}
	g.runGame()
}
