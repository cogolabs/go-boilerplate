package main

import (
	"flag"
	"log"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/cogolabs/go-boilerplate/pkg/web"
	"github.com/cogolabs/terminator"
)

const (
	maxConns = 20
	timeout  = time.Second * 30
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "HTTP bind address")

	shutdownAfter  = flag.Duration("shutdownAfter", 12*time.Hour, "Duration to wait before shutting down (with jitter)")
	shutdownJitter = flag.Duration("shutdownJitter", 2*time.Hour, "Jitter duration in either direction of shutdownAfter")
)

func init() {
	// TODO: set sane HTTP client defaults that make sense for your service
	http.DefaultClient.Timeout = timeout

	http.DefaultTransport.(*http.Transport).MaxConnsPerHost = maxConns
	http.DefaultTransport.(*http.Transport).MaxIdleConns = maxConns
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = maxConns
}

func main() {
	flag.Parse()

	log.Println("starting server")
	defer log.Println("server done")

	srv := &http.Server{
		Addr:         *addr,
		Handler:      web.NewRouter(),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	err := terminator.ServeAndShutdownAfter(&terminator.Options{
		Server:                 srv,
		ShutdownAfter:          *shutdownAfter,
		Jitter:                 *shutdownJitter,
		GracefulShutdownPeriod: time.Minute,
	})
	if err != nil {
		log.Fatalf("error in server: %v\n", err)
	}
}
