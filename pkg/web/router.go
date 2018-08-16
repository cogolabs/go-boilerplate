package web

import (
	"net/http/pprof"

	"github.com/gorilla/mux"
)

// NewRouter instantiates a gorilla mux router and adds all relevant handlers
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)
	r = AddDebug(r)
	return r
}

// AddDebug adds all the /debug/pprof routes to the mux router
func AddDebug(r *mux.Router) *mux.Router {
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	return r
}
