package routines

import cw "TCellConsoleWrapper/tcell_wrapper"

const (
	TITLE_COLOR = cw.DARK_BLUE
	TEXT_COLOR = cw.BEIGE
)

func drawTitle(title string) {
	cw.SetColor(cw.BLACK, TITLE_COLOR)
	cw.PutString("     "+title+"     ", 10, 0)
	cw.SetBgColor(cw.BLACK)
}

func ShowSingleChoiceMenu(title string, lines []string) int { //returns the index of selected line or -1 if nothing was selected.
	val := lines
	cursor := 0
	for {
		cw.Clear_console()
		drawTitle(title)
		for i, _ := range val {
			if cursor == i {
				cw.SetColor(cw.BLACK, TEXT_COLOR)
			} else  {
				cw.SetColor(TEXT_COLOR, cw.BLACK)
			}
			cw.PutString(" "+ val[i] +" ", 1, 1+i)
			cw.SetBgColor(cw.BLACK)
		}
		cw.Flush_console()
		key := cw.ReadKey()
		switch key {
		case "2":
			cursor++
			if cursor == len(val) {
				cursor = 0
			}
		case "8":
			cursor--
			if cursor < 0 {
				cursor = len(val) - 1
			}
		case "ENTER":
			return cursor
		case "ESCAPE":
			return -1
		}
	}
}
