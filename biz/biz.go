package biz

import (
	pb "async_logger/codegen"

	"google.golang.org/grpc"
)

type ServerAPI struct {
	pb.UnimplementedBizServer
}

//
//func Check(context.Context, *pb.Nothing) (*pb.Nothing, error) {
//
//}

func RegisterBizAPI(gRPC *grpc.Server) {
	pb.RegisterBizServer(gRPC, &ServerAPI{})
}
