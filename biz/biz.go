package biz

import (
	pb "async_logger/codegen"
	"context"

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

func (s ServerAPI) Check(ctx context.Context, in *pb.Nothing) (*pb.Nothing, error) {
	return &pb.Nothing{}, nil
}

func (s *ServerAPI) Add(
	_ context.Context, _ *pb.Nothing,
) (*pb.Nothing, error) {
	return &pb.Nothing{}, nil
}

func (s *ServerAPI) Test(
	_ context.Context, _ *pb.Nothing,
) (*pb.Nothing, error) {
	return &pb.Nothing{}, nil
}
