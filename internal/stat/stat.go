package stat

import "time"

type Stat struct {
	subscribers []chan *StatisticsRecord
	byMethod    map[string]uint64
	byConsumer  map[string]uint64
}

type StatisticsRecord struct {
	Timestamp  int64
	ByMethod   map[string]uint64
	ByConsumer map[string]uint64
}

func (s *Stat) Subscribe() chan *StatisticsRecord {
	ch := make(chan *StatisticsRecord, 100)
	s.subscribers = append(s.subscribers, ch)

	return ch
}

func New() *Stat {
	return &Stat{
		subscribers: make([]chan *StatisticsRecord, 0),
		byMethod:    make(map[string]uint64),
		byConsumer:  make(map[string]uint64),
	}
}

// отправлять будем в момент тикера тик

func (s *Stat) sendStatToSubs() {
	// пробегаемся по каналам и шлем событие в каждый из них
	for _, ch := range s.subscribers {
		ch <- &StatisticsRecord{
			Timestamp:  time.Now().Unix(),
			ByConsumer: s.byConsumer,
			ByMethod:   s.byMethod,
		}
	}
}

func (s *Stat) UpdateStat(method string, consumer string) {
	s.byMethod[method]++
	s.byConsumer[consumer]++
	s.sendStatToSubs()
}

// по мапе должны идти подсчеты
// map[method]++
// надо только придумать красивое имя метода
// встраивать будем туда же в интерсептор
// потом по тикеру будем слать всем подписчикам
