package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/04Akaps/Jenkins_docker_go.git/controller"
	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
}

func HttpServerInit() error {
	log.Println(" ------ Server Start ------ ")

	return http.ListenAndServe(":8080", RegisterRouter())
}

func RegisterRouter() http.Handler {
	r := newRouter()
	r.healthCheckRouter()

	return r.router
}

func (r *Router) healthCheckRouter() {
	healthChecker := controller.NewHealthChecker()
	healthCheckRouter := r.router.PathPrefix("/health").Subrouter()
	healthCheckRouter.HandleFunc("", healthChecker.CheckHealth).Methods("GET")
}

func PrintRouters() {
	router := RegisterRouter().(*mux.Router)

	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		methods, _ := route.GetMethods()
		path, _ := route.GetPathTemplate()
		log.Printf("%s: %s\n", strings.Join(methods, ", "), path)
		return nil
	})
	if err != nil {
		fmt.Println("errrrr", err)
	}
}

func newRouter() *Router {
	return &Router{
		router: mux.NewRouter(),
	}
}
