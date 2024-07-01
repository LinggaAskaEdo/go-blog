package rest

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/linggaaskaedo/go-blog/src/business/dto"
	preference "github.com/linggaaskaedo/go-blog/stdlib/preference"
)

func (e *rest) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["userID"]
	e.log.Debug().Str("user_param", userID).Send()

	vid, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		e.httpRespError(w, r, http.StatusBadRequest, err)
		return
	}

	var cacheControl dto.CacheControl
	if r.Header.Get(preference.CacheControl) == preference.CacheMustRevalidate {
		cacheControl.MustRevalidate = true
	} else if r.Header.Get(preference.CacheControl) == "must-db-revalidate" {
		cacheControl.MustDbValidate = true
	}

	result, err := e.uc.User.GetUserByUserID(r.Context(), cacheControl, vid)
	if err != nil {
		e.httpRespError(w, r, http.StatusInternalServerError, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, result, nil)
}
