package main

import (
	"log"

	"github.com/04Akaps/Jenkins_docker_go.git/router"
)

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

var dvs []Device

var version string

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // 시간을 로그로 찍음

	version = "2.10.5"
	dvs = []Device{
		{1, "5F-33-CC-1F-43-82", "2.1.6"},
		{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
	}
}

func main() {
	if err := router.HttpServerInit(); err != nil {
		log.Fatal("Server Start Failed : ", err)
	}
}
