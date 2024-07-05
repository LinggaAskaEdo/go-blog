package router

import (
	"github.com/gorilla/mux"
	"github.com/linggaaskaedo/go-blog/stdlib/middleware"
)

type Options struct {
}

func Init(middleware middleware.Middleware, opt Options) *mux.Router {
	mux := mux.NewRouter()
	mux.Use(middleware.Handler)
	mux.Use(middleware.CORS)

	return mux
}
