package logger

import (
	"time"
)

// не подтягивал целую либу pb а написал свой Ивент

type Event struct {
	Timestamp int64
	Consumer  string
	Method    string
	Host      string
}

type Logger struct {
	log chan *Event
}

func (l *Logger) Log(consumer, method, host string) {
	l.log <- &Event{
		Timestamp: time.Now().Unix(),
		Consumer:  consumer,
		Method:    method,
		Host:      host,
	}
}

func (l *Logger) GetLog() *Event {
	return <-l.log
}

func New() *Logger {
	return &Logger{
		log: make(chan *Event),
	}
}
