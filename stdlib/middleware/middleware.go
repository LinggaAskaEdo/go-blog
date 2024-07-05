package middleware

import (
	"net/http"
	"sync"

	"github.com/rs/zerolog"
)

var once = &sync.Once{}

type Middleware interface {
	Handler(http.Handler) http.Handler
	CORS(http.Handler) http.Handler
}

type middleware struct {
	log zerolog.Logger
}

type Options struct {
}

func Init(log zerolog.Logger, opt Options) Middleware {
	var m *middleware

	once.Do(func() {
		m = &middleware{
			log: log,
		}
	})

	return m
}
