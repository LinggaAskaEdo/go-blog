package router

import (
	"github.com/gorilla/mux"
)

type Options struct {
}

func Init(opt Options) *mux.Router {
	return mux.NewRouter()
}
