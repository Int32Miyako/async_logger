package main

import (
	"context"
	"encoding/json"

	//"errors"
	"fmt"
	"net"
	//pb "async_logger/codegen"
)

/*
тебе нужно сгенерировать grpc сервер
написать для него обвязку и вставить туда вызов сгенерированного кода grpc сервера

добавить реализацию чтобы оно работало по требованиям:
1. ограничение доступа (acl)
2. логирование запросов
3. статистика запросов
*/

type Service struct {
}

// ACL — Access Control List
func parseACL(ACLData string) (map[string][]string, error) {

	acl := make(map[string][]string)
	json.Unmarshal([]byte(ACLData), &acl)

	return acl, nil
}

func StartMyMicroservice(ctx context.Context, listenAddr, ACLData string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer lis.Close()

	_ = &Service{}

	acl, err := parseACL(ACLData)
	if err != nil {
		return err
	}

	fmt.Print(acl)
	return nil
}

func (s *Service) StartTransliteration(ctx context.Context, listenAddr, ACLData string) error {
	return nil
}
