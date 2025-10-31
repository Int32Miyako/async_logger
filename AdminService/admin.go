package AdminService

import (
	"async_logger/AdminService/logging"
	"async_logger/AdminService/statistics"
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

func (s *ServerAPI) Statistics(
	interval *pb.StatInterval,
	server pb.Admin_StatisticsServer,
) error {
	statistics.GetStatistics()
	return nil
}

func (s *ServerAPI) Logging(
	nothing *pb.Nothing,
	server pb.Admin_LoggingServer,

) error {
	logging.GetLogs()
	return nil
}
