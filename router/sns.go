package router

import (
	"github.com/04Akaps/Jenkins_docker_go.git/controller"
)

func (r *Router) SnsRouter() {
	snsController := controller.NewPostController()
	commentController := controller.NewCommentController()
	postPrefixUrl := r.router.PathPrefix("/sns").Subrouter()

	postPrefixUrl.HandleFunc("/getAll/{eoaAddress}", snsController.GetAllPostByEoaAddress).Methods("GET")
	postPrefixUrl.HandleFunc("/{id}", snsController.GetPostByID).Methods("GET")
	postPrefixUrl.HandleFunc("/makeSns", snsController.MakePost).Methods("POST")

	commentPrefixUrl := r.router.PathPrefix("/comment").Subrouter()

	commentPrefixUrl.HandleFunc("/makeComment", commentController.CreateNewComment).Methods("POST")
}
