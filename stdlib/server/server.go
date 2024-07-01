package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var once = &sync.Once{}

type httpServer struct {
	logger  log.Logger
	servers []*http.Server
	opt     Options
}

type Options struct {
	Port         int
	WriteTimeout int
	ReadTimeout  int
	IdleTimeout  int
}

func Init(opt Options, mux *mux.Router) *http.Server {
	var server *http.Server

	once.Do(func() {
		serverPort := fmt.Sprintf(":%d", opt.Port)

		server = &http.Server{
			Addr:         serverPort,
			WriteTimeout: time.Second * time.Duration(opt.WriteTimeout),
			ReadTimeout:  time.Second * time.Duration(opt.ReadTimeout),
			IdleTimeout:  time.Second * time.Duration(opt.IdleTimeout),
			Handler:      mux,
		}
	})

	return server
}
