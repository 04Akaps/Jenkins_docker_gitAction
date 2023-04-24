package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthChecker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(NewHealthChecker().CheckHealth))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but got %d", resp.StatusCode)
	}
}
