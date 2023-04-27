package router

import (
	"github.com/04Akaps/Jenkins_docker_go.git/controller"
)

func (r *Router) healthCheckRouter() {
	healthChecker := controller.NewHealthChecker()
	healthCheckR := r.router.PathPrefix("/health").Subrouter()

	healthCheckR.HandleFunc("", healthChecker.CheckHealth).Methods("GET")
	healthCheckR.HandleFunc("/err", healthChecker.ErrorHealth).Methods("GET")
	healthCheckR.HandleFunc("/body", healthChecker.BodyHealth).Methods("POST")
}
