package rest

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-blog/src/business/usecase"
	"github.com/linggaaskaedo/go-blog/stdlib/middleware"
	"github.com/linggaaskaedo/go-blog/stdlib/parser"
)

var (
	once      = &sync.Once{}
	AppRoutes []RoutePrefix
)

type rest struct {
	log  zerolog.Logger
	mux  *mux.Router
	mw   middleware.Middleware
	json parser.JSONParser
	uc   *usecase.Usecase
}

type Options struct{}

func Init(log zerolog.Logger, mux *mux.Router, mw middleware.Middleware, parse parser.Parser, uc *usecase.Usecase) {
	var e *rest

	once.Do(func() {
		e = &rest{
			log:  log,
			mux:  mux,
			mw:   mw,
			json: parse.JSONParser(),
			uc:   uc,
		}

		e.Serve()
	})
}

func (e *rest) Serve() {
	AppRoutes = e.InitRoute()

	for _, route := range AppRoutes {
		routePrefix := e.mux.PathPrefix(route.Prefix).Subrouter()

		for _, r := range route.SubRoutes {
			var handler http.Handler

			//check to see if route should be protected with jwt
			if r.Protected {
				// TO DO
				// handler = middleware.JWTMiddleware(r.HandlerFunc)
			} else {
				handler = r.HandlerFunc
			}

			routePrefix.
				Path(r.Pattern).
				Handler(handler).
				Methods(r.Method).
				Name(r.Name)
		}
	}
}
