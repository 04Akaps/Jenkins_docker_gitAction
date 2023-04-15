package main

import (
	"log"
	"sync"

	"github.com/04Akaps/Jenkins_docker_go.git/router"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // 시간을 로그로 찍음
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		if err := router.HttpServerInit(); err != nil {
			log.Fatal("Server Start Failed : ", err)
		}
	}()

	go func() {
		defer wg.Done()
		router.PrintRouters()
	}()
}
