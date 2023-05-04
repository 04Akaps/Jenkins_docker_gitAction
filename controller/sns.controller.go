package controller

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"github.com/04Akaps/Jenkins_docker_go.git/crypto"
	connection "github.com/04Akaps/Jenkins_docker_go.git/mysql"
	sqlc "github.com/04Akaps/Jenkins_docker_go.git/mysql/sqlc"
	"github.com/gorilla/mux"
)

type SnsController struct {
	Ctx         context.Context
	MySQLClient *sqlc.Queries
	EthClient   *crypto.CryptoClient
}

type SnsImpl interface {
	GetSnsByID(http.ResponseWriter, *http.Request)
	GetAllSnsByEoaAddress(http.ResponseWriter, *http.Request)
	MakeSns(http.ResponseWriter, *http.Request)
}

var re = regexp.MustCompile("^0x[0-9a-fA-F]{40}$") // 40자리의 16진수 인지 검증

func NewSnsController() SnsImpl {
	context := context.Background()
	endPoint := "https://mainnet.infura.io/v3/299623e5cf3442c8bb2dbe870d8f7d88"
	// 어차피 개인 프로젝트이기 떄문에 Fix

	client := crypto.NewEthClient(context, endPoint)

	return &SnsController{Ctx: context, MySQLClient: connection.NewMySQLClient("sns"), EthClient: client}
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

	if re.MatchString(address) {
		log.Println("16진수가 아니고 40글자도 아닌경우")
		// w.WriteHeader(http.StatusBadGateway)
		// w.Header().Add("error", "잘못된 16진수 또는 40글자가 맞지 않습니다.")
		// return
	}

	w.WriteHeader(http.StatusOK)
}

func (sc *SnsController) MakeSns(w http.ResponseWriter, r *http.Request) {
	log.Println("MakeSns")

	w.WriteHeader(http.StatusOK)
}
