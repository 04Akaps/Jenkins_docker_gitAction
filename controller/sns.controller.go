package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"

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
	GetAllSnsByEoaAddress(http.ResponseWriter, *http.Request)
	MakeSns(http.ResponseWriter, *http.Request)
}

var re = regexp.MustCompile("^0x[0-9a-fA-F]{40}$") // 40자리의 16진수 인지 검증

func NewSnsController() SnsImpl {
	return &SnsController{Ctx: context.Background(), MySQLClient: connection.NewMySQLClient("sns")}
}

func (sc *SnsController) GetSnsByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	log.Println("GetSnsByID", id)

	w.WriteHeader(http.StatusOK)
}

func (sc *SnsController) GetAllSnsByEoaAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["eoaAddress"]

	log.Println("GetAllSnsByEoaAddress", address)

	// 유효한 주소 체크는 나중에

	eoaAddressIsBool := re.MatchString(address)

	if re.MatchString(address) {
		w.WriteHeader(http.StatusBadGateway)
		w.Header().Add("error", "잘못된 16진수 또는 40글자가 맞지 않습니다.")
		return
	}

	fmt.Println(eoaAddressIsBool)

	w.WriteHeader(http.StatusOK)
}

func (sc *SnsController) MakeSns(w http.ResponseWriter, r *http.Request) {
	log.Println("MakeSns")

	w.WriteHeader(http.StatusOK)
}
