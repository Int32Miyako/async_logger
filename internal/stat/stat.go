package stat

type (
	Stat struct {
		subscribers []Subscriber
		IsStarted   bool
		*StatisticsRecord
	}
	Subscriber struct {
		Ch    chan *StatisticsRecord // для транспорта
		State *StatisticsRecord      // хранения
	}
	StatisticsRecord struct {
		Timestamp  int64
		ByMethod   map[string]uint64
		ByConsumer map[string]uint64
	}
)

func (s *Stat) Subscribe() chan *StatisticsRecord {
	ch := make(chan *StatisticsRecord, 100)
	sub := Subscriber{
		Ch: ch,
		State: &StatisticsRecord{
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
		StatisticsRecord: &StatisticsRecord{
			ByMethod:   make(map[string]uint64),
			ByConsumer: make(map[string]uint64),
		},
	}
}

// отправлять будем в момент тикера тик

func (s *Stat) sendStatToSubs() {
	// пробегаемся по каналам и шлем событие в каждый из них
	for _, sub := range s.subscribers {
		sub.Ch <- s.StatisticsRecord
	}
}

func (s *Stat) UpdateStat(method string, consumer string) {
	s.StatisticsRecord.ByMethod[method]++
	s.StatisticsRecord.ByConsumer[consumer]++
	s.sendStatToSubs()
}

func (s *Stat) ResetStat() {
	s.StatisticsRecord = &StatisticsRecord{
		ByMethod:   make(map[string]uint64),
		ByConsumer: make(map[string]uint64),
	}
}

// по мапе должны идти подсчеты
// map[method]++
// надо только придумать красивое имя метода
// встраивать будем туда же в интерсептор
// потом по тикеру будем слать всем подписчикам
