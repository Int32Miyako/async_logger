package stat

import "time"

type (
	Stat struct {
		subscribers []Subscriber
		IsStarted   bool
	}
	Subscriber struct {
		Ch    chan StatisticsRecord // для транспорта
		State StatisticsRecord      // хранения
	}
	StatisticsRecord struct {
		Timestamp  int64
		ByMethod   map[string]uint64
		ByConsumer map[string]uint64
	}
)

func (s *Stat) Subscribe() chan StatisticsRecord {
	ch := make(chan StatisticsRecord, 100)
	sub := Subscriber{
		Ch: ch,
		State: StatisticsRecord{
			ByMethod:   make(map[string]uint64),
			ByConsumer: make(map[string]uint64),
		},
	}

	s.subscribers = append(s.subscribers, sub)

	return ch
}

func New() *Stat {
	return &Stat{
		subscribers: make([]Subscriber, 0),
	}
}

// отправлять будем в момент тикера тик

func (s *Stat) SendStatToSubs() {
	// пробегаемся по каналам и шлем событие в каждый из них
	for i := range s.subscribers {
		s.subscribers[i].Ch <- s.subscribers[i].State
	}
}

func (s *Stat) UpdateStat(method string, consumer string) {
	for i := range s.subscribers {
		s.subscribers[i].State.ByMethod[method]++
		s.subscribers[i].State.ByConsumer[consumer]++
	}
}

func (s *Stat) ResetStat(ch chan StatisticsRecord) {
	for i := range s.subscribers {
		if s.subscribers[i].Ch == ch {
			s.subscribers[i].State.ByMethod = make(map[string]uint64)
			s.subscribers[i].State.ByConsumer = make(map[string]uint64)
			s.subscribers[i].State.Timestamp = time.Now().Unix() // обновляем метку времени
			return                                               // нашли нужного подписчика, дальше не нужно
		}
	}

}

// по мапе должны идти подсчеты
// map[method]++
// надо только придумать красивое имя метода
// встраивать будем туда же в интерсептор
// потом по тикеру будем слать всем подписчикам
