package router

import (
	"github.com/04Akaps/Jenkins_docker_go.git/controller"
)

func (r *Router) SnsRouter() {
	snsController := controller.NewSnsController()
	snsPrefixUrl := r.router.PathPrefix("/sns").Subrouter()

	snsPrefixUrl.HandleFunc("/getAll/{eoaAddress}", snsController.GetAllSnsByEoaAddress).Methods("GET")
	snsPrefixUrl.HandleFunc("/{id}", snsController.GetSnsByID).Methods("GET")
	snsPrefixUrl.HandleFunc("/makeSns", snsController.MakeSns).Methods("POST")
}
