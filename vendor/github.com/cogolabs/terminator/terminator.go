package terminator

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Options struct {
	// Server is the HTTP server that will be started and eventually
	// gracefully shut down
	Server *http.Server

	// ShutdownAfter indicates how long we should keep the HTTP
	// server alive before starting to shutdown
	ShutdownAfter time.Duration

	// Jitter represents a random offset in either direction of
	// ShutdownAfter that we'll add to ShutdownAfter. This ensures
	// that not all servers go down at once, thus avoiding outages.
	Jitter time.Duration

	// GracefulShutdownPeriod represents the maximum amount of time
	// we allow in-flight requests to finish before we force shutdown
	GracefulShutdownPeriod time.Duration
}

// ServeAndShutdownAfter starts the given HTTP server, waits for the given shutdownAfter duration,
// then gracefully shuts the server down and returns to the caller. The maximum amount of time
// in-progress requests are given is represented by gracefulShutdownTimeout.
func ServeAndShutdownAfter(opts *Options) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		// add a random jitter in the range of (-n, n) seconds to the max shutdown
		// time to ensure not all servers in a deployment go down at the same time
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))
		jitterSecs := int(opts.Jitter.Seconds())
		offset := rng.Intn(2*jitterSecs) - jitterSecs
		offsetSecs := time.Second * time.Duration(offset)
		waitFor := opts.ShutdownAfter + offsetSecs

		// wait for a SIGINT to arrive or for the wait duration to elapse
		select {
		case <-sig:
		case <-time.After(waitFor):
		}

		ctx, cancel := context.WithTimeout(context.Background(), opts.GracefulShutdownPeriod)
		defer cancel()

		// gracefully shut down the server
		opts.Server.SetKeepAlivesEnabled(false)
		opts.Server.Shutdown(ctx)
	}()

	err := opts.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	wg.Wait()
	return nil
}
