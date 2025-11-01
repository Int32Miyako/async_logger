package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	FirstTest()
}

func FirstTest() {
	listenAddr := "127.0.0.1:8082"

	ACLData := `{
	"logger1":          ["/main.Admin/Logging"],
	"logger2":          ["/main.Admin/Logging"],
	"stat1":            ["/main.Admin/Statistics"],
	"stat2":            ["/main.Admin/Statistics"],
	"biz_user":         ["/main.Biz/Check", "/main.Biz/Add"],
	"biz_admin":        ["/main.Biz/*"],
	"after_disconnect": ["/main.Biz/Add"]
}`

	ctx, finish := context.WithCancel(context.Background())
	err := StartMyMicroservice(ctx, listenAddr, ACLData)
	if err != nil {
		fmt.Printf("cant start server initial: %v", err)
	}

	mut := &sync.Mutex{}
	mut.Lock()
	finish()
	mut.Unlock()

}

// хочу поотправлять постманом запросы на сервер

func SecondTest() {
	listenAddr := "127.0.0.1:8082"

	ACLData := `{
	"logger1":          ["/main.Admin/Logging"],
	"logger2":          ["/main.Admin/Logging"],
	"stat1":            ["/main.Admin/Statistics"],
	"stat2":            ["/main.Admin/Statistics"],
	"biz_user":         ["/main.Biz/Check", "/main.Biz/Add"],
	"biz_admin":        ["/main.Biz/*"],
	"after_disconnect": ["/main.Biz/Add"]
}`

	ctx, finish := context.WithCancel(context.Background())
	err := StartMyMicroservice(ctx, listenAddr, ACLData)
	if err != nil {
		fmt.Printf("cant start server initial: %v", err)
	}

	mut := &sync.Mutex{}
	mut.Lock()
	finish()
	mut.Unlock()

}

// вот что я понял по тестам
// условно, будет свитч кейс по именам юзеров
// если юзер logger1, то он может дергать только метод Logging
// если юзер stat1, то он может дергать только метод Statistics
// если юзер biz_user, то он может дергать методы Check и Add
// если юзер biz_admin, то он может дергать все методы Biz
// если юзер after_disconnect, то он может дергать метод Add
// все остальные юзеры не могут ничего дергать
// нужно реализовать эту логику в grpc сервере
