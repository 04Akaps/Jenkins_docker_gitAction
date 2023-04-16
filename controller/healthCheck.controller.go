package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type HealthChecker struct {
	healthCtx context.Context
}

func (h *HealthChecker) CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(" --- HealthChecker --- ")

	log.Println("Lint를 위한 단순 err : ", err)
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{healthCtx: context.Background()}
}
