package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	// @TODO: use a specific prometheus namespace to provide unique,
	// fully-qualified metric names
	namespace = "TODO"
)

var (
	// Health indicates whether the server is alive, accepting requests,
	// and is able to successfully operate in a production environment
	Health = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "health",
			Namespace: namespace,
			Help:      "Whether or not the server is healthy",
		},
		[]string{"code"},
	)
)
