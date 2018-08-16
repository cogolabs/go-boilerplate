package web

import "net/http"

// HealthCheckHandler a very simple health check. In production you
// should be running the healthcheck against all external services,
// for example the database, filesystem, cache, etc.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("OK"))
}
