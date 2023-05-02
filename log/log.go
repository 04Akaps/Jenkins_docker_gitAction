package log

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/04Akaps/Jenkins_docker_go.git/monitoring"
	"github.com/prometheus/client_golang/prometheus"
)

const STANDARD_HTTP_ERROR_MESSAGE = "잘못된 요청입니다. 요청을 확인해 주세요"

type errorResponse struct {
	Http_status_code int    `json:"httpStatusCode"`
	Message          string `json:"message"`
	Err_Message      string `json:"errorMessage"`
}

func ServerLogger(next http.Handler, logFile *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "" {
			r.Header.Set("Content-Type", "application/json")
		}
		w.Header().Set("Content-Type", "application/json")

		log.Printf("%s %s", r.Method, r.URL.Path)

		counter, ok := monitoring.RequestCounters[r.URL.Path]

		if !ok {
			// 존재하지 않는 라우터인 경우
			response := &errorResponse{
				Http_status_code: 404,
				Message:          STANDARD_HTTP_ERROR_MESSAGE,
				Err_Message:      "존재하지 않는 라우팅",
			}

			_ = json.NewEncoder(w).Encode(&response)
			return
		}

		counter.With(prometheus.Labels{"type": "router"}).Inc() // Http요청에 대한 counter 증가

		recoder := httptest.NewRecorder()
		next.ServeHTTP(recoder, r)

		if recoder.Code != http.StatusOK {
			// 전송한 요청이 Error에 대한 값을 반환 할 떄
			// 이떄 에러 코드를 기록하고 에러를 사용자에게 반환
			errMessage := recoder.Header().Get("error") // 에러에 대한 정보를 출력

			logFile.Println("--> ", errMessage) // 에러데 대한 메시지를 log파일에 추가

			response := &errorResponse{
				Http_status_code: recoder.Code,
				Message:          STANDARD_HTTP_ERROR_MESSAGE,
				Err_Message:      errMessage,
			}

			_ = json.NewEncoder(w).Encode(&response)
			return
		}
		_, _ = w.Write(recoder.Body.Bytes()) // Body값을 고려하기 위해
	})
}

func GetLogFile(path string) *log.Logger {
	t := time.Now()
	startTime := t.Format("2006-01-02 15:04:05")
	logFile, err := os.Create("log/err/" + startTime + ".log")
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(logFile, "", log.LstdFlags)

	return logger
}
