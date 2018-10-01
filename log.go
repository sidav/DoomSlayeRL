package main

import "fmt"

const LOG_HEIGHT = 5

type LOG struct {
	last_msgs [LOG_HEIGHT]string
}

func (l *LOG) appendMessage(msg string) {
	for i := 0; i < LOG_HEIGHT-1; i++ {
		l.last_msgs[i] = l.last_msgs[i+1]
	}
	l.last_msgs[LOG_HEIGHT-1] = msg
}

func (l *LOG) appendMessagef(msg string, zomg interface{}) {
	for i := 0; i < LOG_HEIGHT-1; i++ {
		l.last_msgs[i] = l.last_msgs[i+1]
	}
	l.last_msgs[LOG_HEIGHT-1] = fmt.Sprintf(msg, zomg)
}

func (l *LOG) warning(msg string) {
	l.appendMessage(msg)
	renderLog(true)
}

func (l *LOG) warningf(msg string, zomg interface{}) {
	l.appendMessagef(msg, zomg)
	renderLog(true)
}
