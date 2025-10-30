package main

import (
	pb "async_logger/codegen"
	"context"
	"encoding/json"

	//"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

/*
тебе нужно сгенерировать grpc сервер
написать для него обвязку и вставить туда вызов сгенерированного кода grpc сервера

добавить реализацию чтобы оно работало по требованиям:
1. ограничение доступа (acl)
2. логирование запросов
3. статистика запросов
*/

func isThisUserExistsInACL(acl map[string][]string, primaryKey string) bool {

	if val, ok := acl[primaryKey]; ok {
		fmt.Print("Users found: ", val)
		return true
	} else {
		fmt.Print("User not found")
		return false

	}
}

type Service struct {
	pb.UnimplementedAdminServer // Это значит:
	// «Я реализую интерфейс BizServer, но всё, чего у меня нет — возьми из UnimplementedAdminServer.»
	pb.UnimplementedBizServer // хз зачем но в доках так
	acl                       map[string][]string
}

// ACL — Access Control List
func parseACL(ACLData string) (map[string][]string, error) {

	acl := make(map[string][]string)
	err := json.Unmarshal([]byte(ACLData), &acl)
	if err != nil {
		return nil, err
	}

	isThisUserExistsInACL(acl, "logger1")

	return acl, nil
}

func StartMyMicroservice(_ context.Context, listenAddr, ACLData string) error {

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer lis.Close()

	service := &Service{}

	s := grpc.NewServer()
	pb.RegisterBizServer(s, service)
	pb.RegisterAdminServer(s, service)
	reflection.Register(s)

	service.Init(ACLData)
	if err != nil {
		return err
	}

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

func (s *Service) Init(ACLData string) error {
	var err error
	s.acl, err = parseACL(ACLData)
	if err != nil {
		return err
	}

	return nil
}
