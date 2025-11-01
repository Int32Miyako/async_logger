package app

import (
	"async_logger/admin"
	"async_logger/biz"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	gRPCServer *grpc.Server
	listenAddr string
}

func New(listenAddr string) *App {
	gRPCServer := grpc.NewServer()

	admin.RegisterServerAPI(gRPCServer)
	biz.RegisterBizAPI(gRPCServer)
	reflection.Register(gRPCServer) // тоже сервис, чтобы посмотреть снаружи наши контракты без proto

	return &App{
		gRPCServer: gRPCServer,
		listenAddr: listenAddr,
	}
}

func (app *App) Start() error {
	lis, err := net.Listen("tcp", app.listenAddr)
	if err != nil {
		return err
	}

	//defer func(lis net.Listener) {
	//	err = lis.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(lis)
	// жизненным циклом управляет graceful shutdown
	// defer lis.Close() не нужен

	return app.gRPCServer.Serve(lis)
}

func (app *App) Stop() error {
	app.gRPCServer.GracefulStop()
	return nil
}
