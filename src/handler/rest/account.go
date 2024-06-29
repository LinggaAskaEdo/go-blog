package rest

import (
	"encoding/json"
	"io"
	"net/http"
)

func (e *rest) AccountLogin(w http.ResponseWriter, r *http.Request) {
	e.logger.Debug().Msg("YUHUUUU")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		e.httpRespError(w, r, http.StatusBadRequest, err)
		return
	}

	var requestBody AccountLoginRequest
	if err := json.Unmarshal(body, &requestBody); err != nil {
		e.httpRespError(w, r, http.StatusBadRequest, err)
		return
	}
}
