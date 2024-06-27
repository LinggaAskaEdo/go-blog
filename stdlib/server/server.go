package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type httpServer struct {
	logger  log.Logger
	servers []*http.Server
	opt     Options
}

type Options struct {
	Port int
}

func Init(logger zerolog.Logger, opt Options, mux *mux.Router) *http.Server {
	serverPort := fmt.Sprintf(":%d", opt.Port)

	server := &http.Server{
		Addr:    serverPort,
		Handler: mux,
	}

	return server
}
