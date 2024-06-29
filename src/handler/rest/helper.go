package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/linggaaskaedo/go-blog/src/business/dto"
	"github.com/linggaaskaedo/go-blog/src/business/entity"
	preference "github.com/linggaaskaedo/go-blog/stdlib/preference"
)

func (e *rest) httpRespSuccess(w http.ResponseWriter, r *http.Request, statusCode int, resp interface{}, p *dto.Pagination) {
	var (
		raw []byte
		err error
	)

	meta := dto.Meta{
		Path:       r.URL.String(),
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    fmt.Sprintf("%s %s [%d] %s", r.Method, r.URL.RequestURI(), statusCode, http.StatusText(statusCode)),
		Error:      "",
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	switch data := resp.(type) {
	case entity.User:
		merchantResp := &HTTPUserResp{
			Meta: meta,
			Data: UserData{
				User: &data,
			},
		}
		raw, err = e.Marshal(merchantResp)
	default:
		e.httpRespError(w, r, http.StatusInternalServerError, errors.New("Invalid response type"))
		return
	}

	if err != nil {
		e.httpRespError(w, r, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set(preference.ContentType, preference.ContentJSON)
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}

func (e *rest) httpRespError(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	log.Error().Stack().Err(err).Send()

	jsonErrResp := &HTTPErrResp{
		Meta: dto.Meta{
			Path:       r.URL.String(),
			StatusCode: statusCode,
			Status:     http.StatusText(statusCode),
			Message:    fmt.Sprintf("%s %s [%d] %s", r.Method, r.URL.RequestURI(), statusCode, http.StatusText(statusCode)),
			Error:      err.Error(),
			Timestamp:  time.Now().Format(time.RFC3339),
		},
	}

	raw, err := json.Marshal(jsonErrResp)
	if err != nil {
		statusCode = http.StatusInternalServerError
		log.Error().Stack().Err(err).Send()
	}

	w.Header().Set(preference.ContentType, preference.ContentJSON)
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}

func (e *rest) Marshal(resp interface{}) ([]byte, error) {
	return json.Marshal(&resp)
}
