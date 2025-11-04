package logger

import "context"

type Logger struct {
	logs chan string

	metrics map[string]int64
}

type ResultLog struct {
	Consumer string
	Method   string
	Host     string
}

func (l *Logger) Log(ctx context.Context) *ResultLog {

	return &ResultLog{
		Consumer: "",
		Method:   "",
		Host:     "",
	}

}

func (l *Logger) GetLog()

func (l *Logger) AddOneToMethodCounter(methodName string) {
	l.metrics[methodName] <- 1
}

func (l *Logger) GetCountOfInvokesMethod(methodName string) int64 {
	l.metrics <- 1
}

func New() *Logger {
	return &Logger{
		logs:    make(chan string),
		metrics: map[string]int64{},
	}
}
