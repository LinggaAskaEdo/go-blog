package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (e *rest) GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	e.logger.Debug().Str("product_param", params["productID"]).Send()
}
