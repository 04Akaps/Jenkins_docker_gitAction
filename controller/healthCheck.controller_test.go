package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthChecker(t *testing.T) {
	log.Println(" ----------- CheckHealth ---------------")

	checkHealthServer := httptest.NewServer(http.HandlerFunc(NewHealthChecker().CheckHealth))
	resp, err := http.Get(checkHealthServer.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but got %d", resp.StatusCode)
	}

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if trimString(HEALTH_CHECK_TEST) != trimString(replaceString(string(response), `"`, "")) {
		t.Errorf("Expected Text is Not Same : %v", string(response))
	}

	log.Println(" ----------- ErrorHealth ---------------")
	// nextRecoder로 넘겨주어도 단순 해당 컨트롤러를 테스트 하는 행위이기 떄문에
	// 상태코드는 controller에 적혀잇는 방식으로 나오는 것을 알아두자!!

	errHealthServer := httptest.NewServer(http.HandlerFunc(NewHealthChecker().ErrorHealth))
	resp, err = http.Get(errHealthServer.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusBadGateway {
		t.Errorf("Expected 502 but got %d", resp.StatusCode)
	}

	response, err = io.ReadAll(resp.Body)

	if err != nil {
		t.Error(err)
	}

	if trimString("") != trimString(string(response)) {
		t.Errorf("Expected Text is Not Same : %v", string(response))
	}

	log.Println(" ----------- Body Health ---------------")

	bodyReq := healthCheckerBodyReq{Name: "test"}
	bodyHealthServer := httptest.NewServer(http.HandlerFunc(NewHealthChecker().BodyHealth))

	bodyReqJSON, err := json.Marshal(bodyReq)
	if err != nil {
		t.Error(err)
	}

	// POST 요청으로 바디 전송
	resp, err = http.Post(bodyHealthServer.URL, "application/json", bytes.NewBuffer(bodyReqJSON))
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200 but got %d", resp.StatusCode)
	}

	var respBody healthCheckerBodyReq
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		t.Error(err)
	}

	// 응답 바디와 요청 바디 일치 여부 확인
	if respBody.Name != bodyReq.Name {
		t.Errorf("Expected body name is %s but got %s", bodyReq.Name, respBody.Name)
	}
}

func replaceString(str, text, replaceText string) string {
	return strings.Replace(str, text, replaceText, -1)
}

func trimString(str string) string {
	return strings.TrimSpace(str)
}
