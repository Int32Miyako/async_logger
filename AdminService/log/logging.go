package log

import (
	"async_logger/AdminService"
	pb "async_logger/codegen"

	"google.golang.org/grpc"
)

type ServerAPI struct {
	AdminService.BaseServer
}

func RegisterServerAPI(grpc *grpc.Server) {
	pb.RegisterAdminServer(grpc, &ServerAPI{})
}

// уже делаем ручки
// на вход должны принять сгенерированный объект запроса
// кроме контекста

func (s *ServerAPI) Logging(
	nothing *pb.Nothing,
	server pb.Admin_LoggingServer,

) error {
	return nil
}
