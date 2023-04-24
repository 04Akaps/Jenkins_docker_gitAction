package main

import (
	"context"
	"log"
	"net/http"

	"github.com/04Akaps/Jenkins_docker_go.git/router"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // 시간을 로그로 찍음
}

// var (
// 	COUNTER = promauto.NewCounter(prometheus.CounterOpts{
// 		Name: "hello_world_total",
// 		Help: "Hello World requested",
// 	})

// 	GAUGE = promauto.NewGauge(prometheus.GaugeOpts{
// 		Name: "hello_world_connection",
// 		Help: "Number of /gauge in progress",
// 	})

// 	SUMMARY = promauto.NewSummary(prometheus.SummaryOpts{
// 		Name: "hello_world_latency_seconds",
// 		Help: "Latency Time for a request /summary",
// 	})

// 	HISTOGRAM = promauto.NewHistogram(prometheus.HistogramOpts{
// 		Name:    "hello_world_latency_histogram",
// 		Help:    "A histogram of Latency Time for a request /histogram",
// 		Buckets: prometheus.LinearBuckets(0.1, 0.1, 10),
// 	})
// )

// func index(w http.ResponseWriter, r *http.Request) {
// 	monitoring.RequestCounters["/health"].Inc()
// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// }

// func gauge(w http.ResponseWriter, r *http.Request) {
// 	GAUGE.Inc()
// 	defer GAUGE.Dec()
// 	time.Sleep(10 * time.Second)
// 	fmt.Fprintf(w, "Gauge, %q", html.EscapeString(r.URL.Path))
// }

// func summary(w http.ResponseWriter, r *http.Request) {
// 	start := time.Now()
// 	defer SUMMARY.Observe(float64(time.Now().Sub(start)))
// 	fmt.Fprintf(w, "Summary, %q", html.EscapeString(r.URL.Path))
// }

// func histogram(w http.ResponseWriter, r *http.Request) {
// 	start := time.Now()
// 	defer HISTOGRAM.Observe(float64(time.Now().Sub(start)))
// 	fmt.Fprintf(w, "Histogram, %q", html.EscapeString(r.URL.Path))
// }

func main() {
	startContext, cancel := context.WithCancel(context.Background())
	reg := prometheus.NewRegistry()

	go func() {
		// HTTP요청 서버
		if err := router.HttpServerInit(reg); err != nil {
			log.Fatal(err)
		}

		defer cancel()
	}()

	go func() {
		// 모니터링 서버

		promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
		http.Handle("/metrics", promHandler)
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatal(err)
		}
		defer cancel()
	}()

	<-startContext.Done()
}
