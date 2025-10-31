package stat

import (
	pb "async_logger/codegen"

	"google.golang.org/grpc"
)

type ServerAPI struct {
	pb.UnimplementedAdminServer
}

func RegisterServerAPI(grpc *grpc.Server) {
	pb.RegisterAdminServer(grpc, &ServerAPI{})
}

func (s *ServerAPI) Statistics(
	interval *pb.StatInterval,
	server pb.Admin_StatisticsServer,
) error {
	return nil
}
