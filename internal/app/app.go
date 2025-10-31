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

	defer func(lis net.Listener) {
		err = lis.Close()
		if err != nil {
			panic(err)
		}
	}(lis)

	err = app.gRPCServer.Serve(lis)
	if err != nil {
		return err
	}

	reflection.Register(app.gRPCServer) // регистрация сервера для рефлексии
	// хз зачем она нужна, но пусть будет
	// и даже это помог написать чатгпт

	return nil
}

func (app *App) Stop() error {
	app.gRPCServer.GracefulStop()
	return nil
}
