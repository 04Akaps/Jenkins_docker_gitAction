package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/04Akaps/Jenkins_docker_go.git/utils"
)

type HealthChecker struct {
	healthCtx context.Context
}

const HEALTH_CHECK_TEST = " --- HealthChecker --- "

func (h *HealthChecker) CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(HEALTH_CHECK_TEST)
}

func (h *HealthChecker) ErrorHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadGateway)
	w.Header().Add("error", "에러에 대해서 작성 - 로그용")
}

type healthCheckerBodyReq struct {
	Name string `json:"name"`
}

func (h HealthChecker) BodyHealth(w http.ResponseWriter, r *http.Request) {
	var req healthCheckerBodyReq

	decoder := utils.BodyDecoder(w, r)
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("Paser Error 체크용", err)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&req)
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{healthCtx: context.Background()}
}
