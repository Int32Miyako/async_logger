package admin

import (
	"async_logger/admin/logging"
	"async_logger/admin/statistics"
	pb "async_logger/codegen"
	"async_logger/internal/logger"

	"google.golang.org/grpc"
)

// уже делаем ручки
// на вход должны принять сгенерированный объект запроса
// кроме контекста

type ServerAPI struct {
	pb.UnimplementedAdminServer
	logger *logger.Logger
}

func RegisterServerAPI(gRPC *grpc.Server, eventLogger *logger.Logger) {
	pb.RegisterAdminServer(gRPC, &ServerAPI{
		logger: eventLogger,
	})
}

func (s *ServerAPI) Logging(
	_ *pb.Nothing,
	server pb.Admin_LoggingServer,
) error {
	err := logging.GetLogs(server, s.logger)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServerAPI) Statistics(
	interval *pb.StatInterval,
	server pb.Admin_StatisticsServer,
) error {
	err := statistics.GetStatistics(interval, server)
	if err != nil {
		return err
	}
	return nil
}
