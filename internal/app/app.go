package app

import (
	"async_logger/admin"
	"async_logger/biz"
	"async_logger/internal/interceptors"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	gRPCServer *grpc.Server
	listenAddr string
	acl        map[string][]string
}

func New(listenAddr string, ACLData map[string][]string) *App {
	gRPCServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.AclInterceptor),
	)

	admin.RegisterServerAPI(gRPCServer)
	biz.RegisterBizAPI(gRPCServer)
	reflection.Register(gRPCServer) // тоже сервис, чтобы посмотреть снаружи наши контракты без proto

	return &App{
		gRPCServer: gRPCServer,
		listenAddr: listenAddr,
		acl:        ACLData,
	}
}

func (app *App) Start() error {
	lis, err := net.Listen("tcp", app.listenAddr)
	if err != nil {
		return err
	}

	return app.gRPCServer.Serve(lis)
}

func (app *App) Stop() error {
	app.gRPCServer.GracefulStop()
	return nil
}
