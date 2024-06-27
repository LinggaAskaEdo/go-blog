package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (e *rest) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	e.logger.Debug().Str("params", params["userId"]).Send()
}
