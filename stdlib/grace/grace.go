package grace

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog"
)

var wg sync.WaitGroup

type App interface {
	Serve()
	Stop()
}

type app struct {
	logger     zerolog.Logger
	httpServer *http.Server
}

type Options struct {
}

func Init(logger zerolog.Logger, httpServer *http.Server) App {
	gs := &app{
		logger:     logger,
		httpServer: httpServer,
	}

	return gs
}

func (g *app) Serve() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go startHTTPServer(ctx, &wg, g.logger, g.httpServer)

	// Listen for termination signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Wait for termination signal
	<-signalCh

	// Start the graceful shutdown process
	g.logger.Debug().Msg("Gracefully shutting down HTTP server...")

	// Cancel the context to signal the HTTP server to stop
	cancel()

	// Wait for the HTTP server to finish
	wg.Wait()

	g.logger.Debug().Msg("Shutdown complete.")
}

func (g *app) Stop() {
	// TODO
}
