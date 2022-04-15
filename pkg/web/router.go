package web

import (
	"net/http"

	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewRouter instantiates a gorilla mux router and adds all relevant handlers
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// add pprof routes
	r.PathPrefix("/debug").Handler(http.DefaultServeMux)

	// add promethus handlers
	r.Handle("/metrics", promhttp.Handler())

	// health check handlers
	r.HandleFunc("/ping", PingHandler)
	r.HandleFunc("/health", HealthHandler)

	// @TODO: Add your routes here...
	// r.HandleFunc("/foo", FooHandler)

	return r
}
