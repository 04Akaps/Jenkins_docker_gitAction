package monitoring

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Metrics struct {
	HttpCounter map[string]prometheus.Counter
}

var RequestCounters = &Metrics{
	HttpCounter: make(map[string]prometheus.Counter),
}

func RegisterMetrics(path string) {
	metricName := strings.Replace(path, "/", "_", -1)

	counter := promauto.NewCounter(
		prometheus.CounterOpts{
			Namespace: "http",
			Subsystem: "requests",
			Name:      metricName,
			Help:      "Number of HTTP requests",
		},
	)

	RequestCounters.HttpCounter[metricName] = counter
}

func MetrisMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metricName := strings.Replace(r.URL.Path, "/", "_", -1)
		counter, ok := RequestCounters.HttpCounter[metricName]
		fmt.Println("Counter", counter)
		log.Println(metricName)
		if ok {
			counter.Inc()
		} else {
			panic("뭔가 잘못된 --- 디버깅용 panic")
		}
		next.ServeHTTP(w, r)
	})
}
