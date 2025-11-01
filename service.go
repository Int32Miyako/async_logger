package main

import (
	"async_logger/internal/app"
	"context"
)

/*
тебе нужно сгенерировать grpc сервер
написать для него обвязку и вставить туда вызов сгенерированного кода grpc сервера

добавить реализацию чтобы оно работало по требованиям:
1. ограничение доступа (acl)
2. логирование запросов
3. статистика запросов
*/

func StartMyMicroservice(ctx context.Context, listenAddr, ACLData string) error {
	service := app.New(listenAddr)

	errCh := make(chan error)
	go func() {
		if err := service.Start(); err != nil {
			errCh <- err
		}
		close(errCh)
	}()

	select {
	case <-ctx.Done():
		return service.Stop()
	case err := <-errCh:
		return err
	}

}
