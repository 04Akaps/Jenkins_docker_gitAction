package controller

import (
	"context"
	"encoding/json"
	"net/http"
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

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{healthCtx: context.Background()}
}
