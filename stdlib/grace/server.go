package grace

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

func startHTTPServer(ctx context.Context, wg *sync.WaitGroup, logger zerolog.Logger, httpServer *http.Server) {
	defer wg.Done()

	// Start the HTTP server in a separate goroutine
	go func() {
		logger.Debug().Msg("Starting HTTP server...")

		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Debug().AnErr("HTTP server error", err)
		}
	}()

	select {
	case <-ctx.Done():
		// Shutdown the server gracefully
		logger.Debug().Msg("Shutting down HTTP server gracefully...")

		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			logger.Debug().AnErr("HTTP server shutdown error", err)
		}
	default:
		// TODO
	}

	logger.Debug().Msg("HTTP server stopped")
}
