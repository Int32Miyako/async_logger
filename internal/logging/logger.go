package logging

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
	subscribers []chan *Event
}

func (l *Logger) Log(consumer, method, host string) {
	// пробегаемся по каналам и шлем событие в каждый из них
	for _, ch := range l.subscribers {
		ch <- &Event{
			Timestamp: time.Now().Unix(),
			Consumer:  consumer,
			Method:    method,
			Host:      host,
		}
	}
}

func (l *Logger) Subscribe() chan *Event {
	// возвращаем последний канал как новый подписчик
	ch := make(chan *Event, 100)
	l.subscribers = append(l.subscribers, ch)

	return ch
}

func New() *Logger {
	// иници
	return &Logger{
		subscribers: make([]chan *Event, 0),
	}
}
