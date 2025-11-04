package main

import (
	"async_logger/internal/acl"
	"async_logger/internal/app"
	"context"
	"errors"
	"fmt"
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
	ACL, err := acl.ParseACL(ACLData)
	if err != nil {
		return errors.New(fmt.Sprintf("ACL parse error, reason: %s %s", err, ACL))
	}

	service := app.New(listenAddr, ACL)

	errCh := make(chan error)
	go func() {
		if err = service.Start(); err != nil {
			errCh <- err
		}
	}()

	go func() {
		<-ctx.Done()
		_ = service.Stop()

	}()

	return nil
}
