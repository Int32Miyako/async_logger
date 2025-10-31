package main

import (
	"async_logger/internal/app"
	"context"

	"fmt"

	"async_logger/internal/acl"
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

	if err := service.Start(); err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		if err := service.Stop(); err != nil {
			return
		}
	}()

	aclData, err := acl.ParseACL(ACLData)
	if err != nil {
		return err
	}
	fmt.Print(aclData)

	return nil
}
