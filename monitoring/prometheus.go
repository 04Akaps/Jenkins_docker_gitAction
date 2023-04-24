package monitoring

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var RequestCounters = make(map[string]*prometheus.CounterVec)

func RegisterMetrics(path string, reg *prometheus.Registry) {
	routerPath := strings.Replace(path, "/", "_", -1)

	counter := promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "go_server",
		Name:      fmt.Sprintf("router%s_count", routerPath),
		Help:      fmt.Sprintf("router%s_Request", routerPath),
	}, []string{"type"})

	reg.MustRegister(counter)

	RequestCounters[path] = counter
}
