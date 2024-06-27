package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

var AppRoutes []RoutePrefix

type RoutePrefix struct {
	Prefix    string
	SubRoutes []Route
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}

type Options struct {
}

func Init(logger zerolog.Logger, opt Options) *mux.Router {
	return mux.NewRouter()
}
