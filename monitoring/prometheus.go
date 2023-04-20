package monitoring

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var RequestCounters = make(map[string]prometheus.Counter)

func RegisterMetrics(path string) {
	routerPath := strings.Replace(path, "/", "_", -1)

	counter := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "go_server",
		Name:      fmt.Sprintf("router%s_count", routerPath),
		Help:      fmt.Sprintf("router%s_Request", routerPath),
	})

	RequestCounters[path] = counter
}
