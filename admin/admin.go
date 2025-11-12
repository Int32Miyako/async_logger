package admin

import (
	pb "async_logger/codegen"
	"async_logger/internal/logging"
	"io"
	"time"

	"google.golang.org/grpc"
)

// уже делаем ручки
// на вход должны принять сгенерированный объект запроса
// кроме контекста

type ServerAPI struct {
	pb.UnimplementedAdminServer
	logger *logging.Logger
}

func RegisterServerAPI(gRPC *grpc.Server, eventLogger *logging.Logger) {
	pb.RegisterAdminServer(gRPC, &ServerAPI{
		logger: eventLogger,
	})
}

func (s *ServerAPI) Logging(
	_ *pb.Nothing,
	server pb.Admin_LoggingServer,
) error {
	logger := s.logger
	ch := logger.Subscribe()

	for {
		event := <-ch

		err := server.Send(&pb.Event{
			Timestamp: event.Timestamp,
			Consumer:  event.Consumer,
			Method:    event.Method,
			Host:      event.Host,
		})

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

func (s *ServerAPI) Statistics(
	interval *pb.StatInterval,
	server pb.Admin_StatisticsServer,
) error {
	ticker := time.NewTicker(time.Duration(interval.GetIntervalSeconds()) * time.Second)
	// ticker генерирует события через фиксированные интервалы времени
	defer ticker.Stop()

	s.logger.Stat.IsStarted = true
	ch := s.logger.Stat.Subscribe()

	for {
		select {
		case <-ticker.C:
			s.logger.Stat.SendStatToSubs()
			stat := <-ch
			err := server.Send(&pb.Stat{
				Timestamp:  stat.Timestamp,
				ByMethod:   stat.ByMethod,
				ByConsumer: stat.ByConsumer,
			})
			s.logger.Stat.ResetStat(ch)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

		case <-server.Context().Done():
			return nil
		}
	}
}
