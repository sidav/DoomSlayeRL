package main

import "GoConsoleWrapper/console_wrapper"

func playerControl(p *pawn) {
	key_pressed := console_wrapper.Read_key_char()
	if key_pressed == 's' {
		p.y += 1
	}
}
