package router

import (
	"log"
	"net/http"

	"github.com/04Akaps/Jenkins_docker_go.git/controller"
	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func HttpServerInit() error {
	log.Println(" ------ Server Start ------ ")

	return http.ListenAndServe(":80", registerRouter())
}

func registerRouter() http.Handler {
	r := newRouter()
	r.healthCheckRouter()

	return r.router
}

func (r *Router) healthCheckRouter() {
	healthChecker := controller.NewHealthChecker()
	healthCheckRouter := r.router.PathPrefix("/health").Subrouter()
	healthCheckRouter.HandleFunc("", healthChecker.CheckHealth).Methods("GET")
}

func newRouter() *Router {
	return &Router{
		router: mux.NewRouter(),
	}
}
