package router

import (
	"github.com/04Akaps/Jenkins_docker_go.git/controller"
)

func (r *Router) healthCheckRouter() {
	healthChecker := controller.NewHealthChecker()
	healthCheckR := r.router.PathPrefix("/health").Subrouter()

	// 일단 Metrics등록을 자동화 하는 방법이 딲히 생각이 안나서 Test하기 위해서 Fix해둔 코드

	healthCheckR.HandleFunc("", healthChecker.CheckHealth).Methods("GET")
	healthCheckR.HandleFunc("/err", healthChecker.ErrorHealth).Methods("GET")
}
