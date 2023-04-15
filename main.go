package main

import (
	"log"

	"github.com/04Akaps/Jenkins_docker_go.git/router"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // 시간을 로그로 찍음
}

func main() {
	router.HttpServerInit()
}
