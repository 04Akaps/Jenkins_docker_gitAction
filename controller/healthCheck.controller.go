package controller

import (
	"context"
	"log"
	"net/http"
)

type HealthChecker struct {
	healthCtx context.Context
}

func (h *HealthChecker) CheckHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("HealthChecker")
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{healthCtx: context.Background()}
}
