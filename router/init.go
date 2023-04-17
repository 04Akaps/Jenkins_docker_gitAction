package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/04Akaps/Jenkins_docker_go.git/controller"
	logger "github.com/04Akaps/Jenkins_docker_go.git/log"
	"github.com/gorilla/mux"
)

type Router struct {
	router  *mux.Router
	logFile *log.Logger
}

func HttpServerInit() error {
	log.Println(" ------ Server Start ------ ")

	r, _ := RegisterRouter()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		// 서버 시작 시 모든 router를 확인하기 위한 sync
		defer wg.Done()
		PrintRouters()
	}()

	wg.Wait()

	return http.ListenAndServe(":8080", r)
}

func RegisterRouter() (http.Handler, *mux.Router) {
	r := newRouter()

	logMux := logger.ServerLogger(r.router, r.logFile)

	r.healthCheckRouter()

	return logMux, r.router
}

func (r *Router) healthCheckRouter() {
	healthChecker := controller.NewHealthChecker()
	healthCheckRouter := r.router.PathPrefix("/health").Subrouter()
	healthCheckRouter.HandleFunc("", healthChecker.CheckHealth).Methods("GET")
	healthCheckRouter.HandleFunc("/err", healthChecker.ErrorHealth).Methods("GET")
}

func PrintRouters() {
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
