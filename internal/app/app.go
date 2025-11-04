package app

import (
	"async_logger/admin"
	"async_logger/biz"
	"async_logger/internal/interceptors"
	"async_logger/internal/logger"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	gRPCServer *grpc.Server
	listenAddr string
	acl        map[string][]string
	logger     *logger.Logger
}

func New(listenAddr string, ACLData map[string][]string) *App {
	eventLogger := logger.New()
	gRPCServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.AclInterceptor(ACLData, eventLogger)),
		grpc.StreamInterceptor(interceptors.AclStreamInterceptor(ACLData, eventLogger)),
	)

	admin.RegisterServerAPI(gRPCServer, eventLogger)
	biz.RegisterBizAPI(gRPCServer)
	reflection.Register(gRPCServer) // тоже сервис, чтобы посмотреть снаружи наши контракты без proto

	return &App{
		gRPCServer: gRPCServer,
		listenAddr: listenAddr,
		acl:        ACLData,
		logger:     eventLogger,
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
