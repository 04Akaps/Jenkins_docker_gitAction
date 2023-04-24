package controller

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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

	// _ = json.NewEncoder(w).Encode(" --- HealthChecker --- ")

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if trimString(HEALTH_CHECK_TEST) != trimString(replaceString(string(response), `"`, "")) {
		t.Errorf("Expected Text is Not Same : %v", string(response))
	}
}

func replaceString(str, text, replaceText string) string {
	return strings.Replace(str, text, replaceText, -1)
}

func trimString(str string) string {
	return strings.TrimSpace(str)
}
