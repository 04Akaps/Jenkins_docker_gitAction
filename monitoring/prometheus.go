package monitoring

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	Devices prometheus.Gauge
	Info    *prometheus.GaugeVec
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		Devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "goServer",
			Name:      "connected_devices",
			Help:      "Number of currently connected devices.",
		}),
		Info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "goServer",
			Name:      "info",
			Help:      "Information about the My App environment.",
		},
			[]string{"version"}),
	}
	reg.MustRegister(m.Devices, m.Info)
	return m
}
