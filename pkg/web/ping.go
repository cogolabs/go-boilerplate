package web

import (
	"fmt"
	"net/http"
)

// PingHandler always returns 200 OK and is used to test whether or
// not the server is accepting requests
func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

// HealthHandler is a very simple health check. In production you
// should be running the healthcheck against all external services,
// for example the database, filesystem, cache, etc.
func HealthHandler(w http.ResponseWriter, r *http.Request) {

	// @TODO: Check all downstream services that are required for
	// your server to function

	w.Write([]byte("OK\n"))
}
