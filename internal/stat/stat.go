package stat

import "time"

type Stat struct {
	subscribers []chan *StatRecord
}

type StatRecord struct {
	Timestamp  int64
	ByMethod   map[string]uint64
	ByConsumer map[string]uint64
}

func (s *Stat) Subscribe() chan *StatRecord {
	ch := make(chan *StatRecord, 100)
	s.subscribers = append(s.subscribers, ch)

	return ch
}

func New() *Stat {
	return &Stat{
		subscribers: make([]chan *StatRecord, 0),
	}
}

// отправлять будем в момент тикера тик

func (l *Stat) SendStatToSubs(consumer, method map[string]uint64) {
	// пробегаемся по каналам и шлем событие в каждый из них
	for _, ch := range l.subscribers {
		ch <- &StatRecord{
			Timestamp:  time.Now().Unix(),
			ByConsumer: consumer,
			ByMethod:   method,
		}
	}
}

func (l *Stat) AddMethodToListenInStatistics() chan map[string]uint64 {
	result := make(chan map[string]uint64)
	return result
}

func (l *Stat) AddConsumerToListenInStatistics() chan map[string]uint64 {
	return nil
}

func (l *Stat) HandleStat(method string, consumer string) {
	// тут нужно собрать статистику и разослать ее подписчикам

}

// по мапе должны идти подсчеты
// map[method]++
// надо только придумать красивое имя метода
// встраивать будем туда же в интерсептор
// потом по тикеру будем слать всем подписчикам
