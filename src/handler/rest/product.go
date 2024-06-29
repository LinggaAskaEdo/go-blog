package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func (e *rest) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	log.Debug().Str("product_param", params["productID"]).Send()
}
