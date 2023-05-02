package controller

import (
	"context"
	"log"
	"net/http"

	connection "github.com/04Akaps/Jenkins_docker_go.git/mysql"
	sqlc "github.com/04Akaps/Jenkins_docker_go.git/mysql/sqlc"
	"github.com/gorilla/mux"
)

type SnsController struct {
	Ctx         context.Context
	MySQLClient *sqlc.Queries
}

type SnsImpl interface {
	GetSnsByID(http.ResponseWriter, *http.Request)
	GetAllSnsByUserName(http.ResponseWriter, *http.Request)
	MakeSns(http.ResponseWriter, *http.Request)
}

func NewSnsController() SnsImpl {
	return &SnsController{Ctx: context.Background(), MySQLClient: connection.NewMySQLClient("sns")}
}

func (sc *SnsController) GetSnsByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("GetSnsByID", id)

	w.WriteHeader(http.StatusOK)
}

func (sc *SnsController) GetAllSnsByUserName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["userName"]

	log.Println("GetAllSnsByUserName", name)

	w.WriteHeader(http.StatusOK)
}

func (sc *SnsController) MakeSns(w http.ResponseWriter, r *http.Request) {
	log.Println("MakeSns")

	w.WriteHeader(http.StatusOK)
}
