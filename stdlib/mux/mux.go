package router

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)


type Options struct {
}

func Init(logger zerolog.Logger, opt Options) *mux.Router {
	return mux.NewRouter()
}
