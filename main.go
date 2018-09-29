package main

import (
	cw "GoConsoleWrapper/console_wrapper"
	"time"
)

func main() {
	cw.Init_console()
	defer cw.Close_console()
	go func() {
		cw.Run_event_listener()
	}()

	g := game{}
	g.runGame()
	time.Sleep(1000 * time.Millisecond)
}
