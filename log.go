package main

import (
	"fmt"
	"strings"
)

const LOG_HEIGHT = 5

var log_berserkify = false

type logMessage struct {
	message string
	count int
	color byte
}

func (m *logMessage) getText() string {
	if m.count > 1 {
		return fmt.Sprintf("%s (x%d)", m.message, m.count)
	} else {
		return m.message
	}
}

type LOG struct {
	last_msgs [LOG_HEIGHT]logMessage
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
	if l.last_msgs[LOG_HEIGHT-1].message == msg {
		l.last_msgs[LOG_HEIGHT-1].count++
	} else {
		for i := 0; i < LOG_HEIGHT-1; i++ {
			l.last_msgs[i] = l.last_msgs[i+1]
		}
		l.last_msgs[LOG_HEIGHT-1] = logMessage{message: msg, count:1}
	}
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
