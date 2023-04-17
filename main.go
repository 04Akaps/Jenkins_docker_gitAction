package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/04Akaps/Jenkins_docker_go.git/monitoring"
	"github.com/04Akaps/Jenkins_docker_go.git/router"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	go func() {
		if err := router.HttpServerInit(); err != nil {
			log.Fatal("Server Start Failed : ", err)
		}
	}()

	reg := prometheus.NewRegistry()
	m := monitoring.NewMetrics(reg)

	m.Devices.Set(float64(len(dvs)))
	m.Info.With(prometheus.Labels{"version": version}).Set(1)

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})

	http.Handle("/metrics", promHandler)
	http.HandleFunc("/devices", getDevices)
	http.ListenAndServe(":8081", nil)
}

func getDevices(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(dvs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
