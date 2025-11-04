package admin

import (
	"async_logger/admin/logging"
	"async_logger/admin/statistics"
	pb "async_logger/codegen"

	"google.golang.org/grpc"
)

// уже делаем ручки
// на вход должны принять сгенерированный объект запроса
// кроме контекста

type ServerAPI struct {
	pb.UnimplementedAdminServer
}

func RegisterServerAPI(gRPC *grpc.Server) {
	pb.RegisterAdminServer(gRPC, &ServerAPI{})
}

func (s *ServerAPI) Logging(
	_ *pb.Nothing,
	server pb.Admin_LoggingServer,
) error {
	err := logging.GetLogs(server)
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
