package observe

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// FakeMetric is an example metric for using prometheus
	FakeMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fake_metric",
			Help: "A fake counter for showing how to setup prometheus",
		},
		[]string{},
	)
)

// RegisterPrometheus adds the prometheus handler to the mux router
// Note you must register every metric with prometheus for it show up
// when the /metrics route is hit.
func RegisterPrometheus(m *mux.Router) *mux.Router {
	prometheus.MustRegister(FakeMetric)

	m.Handle("/metrics", promhttp.Handler())
	return m
}
