package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	logger "github.com/04Akaps/Jenkins_docker_go.git/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Router struct {
	router  *mux.Router
	logFile *log.Logger
}

func HttpServerInit() error {
	log.Println(" ------ Server Start ------ ")

	logMux, _ := RegisterRouter()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// 서버 시작 시 모든 router를 확인하기 위한 sync
		defer wg.Done()
		// printRouters()
	}()

	wg.Wait()

	return http.ListenAndServe(":8080", logMux)
}

func RegisterRouter() (http.Handler, *mux.Router) {
	r := newRouter()

	logMux := logger.ServerLogger(r.router, r.logFile)

	reg := prometheus.NewRegistry()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	r.router.Handle("/metrics", promHandler)
	// reg.MustRegister를 사용해야 데이터가 보이는지 체크 필요
	r.healthCheckRouter()

	return logMux, r.router
}

// ---- Utils Function ----

func printRouters() {
	_, router := RegisterRouter()

	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, _ := route.GetMethods()
		path, _ := route.GetPathTemplate()

		if methods != nil {
			log.Printf("%s: %s\n", strings.Join(methods, ", "), path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("errrrr", err)
	}
}

func newRouter() *Router {
	logFile := logger.GetLogFile(".")
	return &Router{
		router:  mux.NewRouter(),
		logFile: logFile,
	}
}
