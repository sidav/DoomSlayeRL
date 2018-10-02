package main

import (
	"fmt"
	"strings"
)

const LOG_HEIGHT = 5

var log_berserkify = false

type LOG struct {
	last_msgs [LOG_HEIGHT]string
}

func (l *LOG) appendMessage(msg string) {
	if log_berserkify {
		msg = strings.ToUpper(msg)
		lastchar := msg[len(msg)-1]
		switch lastchar {
		case '?':
			msg = msg[0:len(msg)-1] + "???"
		case '!':
			msg = msg[0:len(msg)-1] + "!!!"
		case '.':
			msg = msg[0:len(msg)-1] + "!!!"
		default:
			msg += "!!!"
		}
	}
	for i := 0; i < LOG_HEIGHT-1; i++ {
		l.last_msgs[i] = l.last_msgs[i+1]
	}
	l.last_msgs[LOG_HEIGHT-1] = msg
}

func (l *LOG) appendMessagef(msg string, zomg interface{}) {
	msg = fmt.Sprintf(msg, zomg)
	l.appendMessage(msg)
}

func (l *LOG) warning(msg string) {
	l.appendMessage(msg)
	renderLog(true)
}

func (l *LOG) warningf(msg string, zomg interface{}) {
	l.appendMessagef(msg, zomg)
	renderLog(true)
}
