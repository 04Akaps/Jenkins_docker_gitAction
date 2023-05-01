package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	logger "github.com/04Akaps/Jenkins_docker_go.git/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

type Router struct {
	router  *mux.Router
	logFile *log.Logger
	reg     *prometheus.Registry
}

type RouterInterface interface {
	registerRouter() (http.Handler, *mux.Router)
	printRouters()
}

func HttpServerInit(reg *prometheus.Registry) error {
	log.Println(" ------ Server Start ------ ")

	// logMux, _ := RegisterRouter()
	r := newRouter(reg)

	logMux, _ := r.registerRouter()
	r.printRouters()

	return http.ListenAndServe(":8080", logMux)
}

func (r Router) registerRouter() (http.Handler, *mux.Router) {
	logMux := logger.ServerLogger(r.router, r.logFile)

	r.healthCheckRouter()
	r.SnsRouter()

	return logMux, r.router
}

func (r Router) printRouters() {
	err := r.router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, _ := route.GetMethods()
		path, _ := route.GetPathTemplate()

		if methods != nil {
			// monitoring.RegisterMetrics(path, r.reg)
			log.Printf("%s: %s\n", strings.Join(methods, ", "), path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("errrrr", err)
	}
}

func newRouter(reg *prometheus.Registry) RouterInterface {
	logFile := logger.GetLogFile(".")
	return &Router{
		router:  mux.NewRouter(),
		logFile: logFile,
		reg:     reg,
	}
}
