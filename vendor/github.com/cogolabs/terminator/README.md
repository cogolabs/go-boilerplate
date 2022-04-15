# Terminator

<img src="https://user-images.githubusercontent.com/8205547/162049312-28e505f3-100c-47b4-a168-78d4ff9dd19f.jpg" />

Library to automatically run and kill HTTP servers at regular intervals. Includes a jitter on the kill interval to ensure that a group of server processes started at roughly the same time don't all go down at once.

## Usage:

```go
package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/cogolabs/terminator"
)

const (
	timeout = 30 * time.Second
)

var (
	shutdownAfter  = flag.Duration("shutdownAfter", 12*time.Hour, "Duration to wait before shutting down (with jitter)")
	shutdownJitter = flag.Duration("shutdownJitter", 2*time.Hour, "Jitter duration in either direction of shutdownAfter")
)

func main() {
	flag.Parse()

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      nil, // your handler here
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	log.Println("starting server")
	err := terminator.ServeAndShutdownAfter(&terminator.Options{
		Server:                 srv,
		ShutdownAfter:          *shutdownAfter,
		Jitter:                 *shutdownJitter,
		GracefulShutdownPeriod: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("error in server: %v\n", err)
	}

	log.Println("server done")
}
```
