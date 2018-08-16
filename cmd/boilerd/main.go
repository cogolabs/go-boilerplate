package main

import (
	"flag"
	"time"

	"net/http"
	_ "net/http/pprof"

	"git.mycompany.com/platform/go-boilerplate.git/pkg/observe"
	"git.mycompany.com/platform/go-boilerplate.git/pkg/web"
	"github.com/facebookgo/grace/gracehttp"
)

var (
	addr  = flag.String("addr", "0.0.0.0:9090", "Primary HTTP addr")
	debug = flag.Bool("debug", false, "Set the logging level to debug")
	dsn   = flag.String("dsn", "", "Set the Raven DSN for sentry alerting")

	writeTimeout = time.Second * 30
	readTimeout  = time.Second * 30
)

func main() {
	flag.Parse()
	observe.InitLogging(*debug, *dsn)
	r := web.NewRouter()
	r = observe.RegisterPrometheus(r)
	srv := &http.Server{
		Handler:      r,
		Addr:         *addr,
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}
	gracehttp.Serve(srv)
}
