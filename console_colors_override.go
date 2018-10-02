package main

import (
	cw "TCellConsoleWrapper/tcell_wrapper"
)

// This file is NOT a console wrapper.
// Its purpose is to override color output for effects (i.g. berserk or invulnerability) creation.

var (
	OVERRIDE_COLOR           bool
	OVERRIDING_COLOR         int
	INVERSE_COLOR            bool
	INVERSE_BLACKWHITE_COLOR bool
)

func setColor(fg, bg int) {
	if OVERRIDE_COLOR {
		cw.SetColor(OVERRIDING_COLOR, cw.BLACK)
		return
	}
	if INVERSE_BLACKWHITE_COLOR {
		cw.SetColor(cw.BLACK, cw.BEIGE)
		return
	}
	cw.SetColor(fg, bg)
}

func setFgColor(fg int) {
	if OVERRIDE_COLOR {
		cw.SetColor(OVERRIDING_COLOR, cw.BLACK)
		return
	}
	if INVERSE_BLACKWHITE_COLOR {
		cw.SetColor(cw.BLACK, cw.BEIGE)
		return
	}
	cw.SetFgColor(fg)
}

func setBgColor(bg int) {
	if OVERRIDE_COLOR {
		cw.SetColor(OVERRIDING_COLOR, cw.BLACK)
		return
	}
	if INVERSE_BLACKWHITE_COLOR {
		cw.SetColor(cw.BLACK, cw.BEIGE)
		return
	}
	cw.SetBgColor(bg)
}

func clearConsole() {
	if INVERSE_BLACKWHITE_COLOR {
		return
	}
	cw.Clear_console()
}
