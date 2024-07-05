package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/linggaaskaedo/go-blog/src/business/dto"
	apperr "github.com/linggaaskaedo/go-blog/stdlib/errors"
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
		Error:      nil,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	switch data := resp.(type) {
	case nil:
		httpResp := &HTTPEmptyResp{
			Meta: meta,
		}
		raw, err = e.Marshal(httpResp)

	case dto.UserDTO:
		userResp := &HTTPUserResp{
			Meta: meta,
			Data: UserData{
				User: &data,
			},
		}
		raw, err = e.Marshal(userResp)

	case dto.DivisionDTO:
		divisionResp := &HTTPDivisionResp{
			Meta: meta,
			Data: DivisionData{
				Division: &data,
			},
		}
		raw, err = e.Marshal(divisionResp)

	default:
		e.httpRespError(w, r, errors.New("Invalid response type"))
		return
	}

	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	w.Header().Set(preference.ContentType, preference.ContentJSON)
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}

func (e *rest) httpRespError(w http.ResponseWriter, r *http.Request, err error) {
	e.log.Error().Stack().Err(err).Send()

	lang := preference.LangID
	if r.Header.Get(preference.AppLang) == preference.LangEN {
		lang = preference.LangEN
	}

	statusCode, displayError := apperr.Compile(apperr.COMMON, err, lang, true)
	statusStr := http.StatusText(statusCode)

	jsonErrResp := &HTTPErrResp{
		Meta: dto.Meta{
			Path:       r.URL.String(),
			StatusCode: statusCode,
			Status:     statusStr,
			Message:    fmt.Sprintf("%s %s [%d] %s", r.Method, r.URL.RequestURI(), statusCode, http.StatusText(statusCode)),
			Error:      &displayError,
			Timestamp:  time.Now().Format(time.RFC3339),
		},
	}

	raw, err := e.json.Marshal(jsonErrResp)
	if err != nil {
		statusCode = http.StatusInternalServerError
		log.Error().Stack().Err(err).Send()
	}

	w.Header().Set(preference.ContentType, preference.ContentJSON)
	w.WriteHeader(statusCode)
	_, _ = w.Write(raw)
}

func (e *rest) Marshal(resp interface{}) ([]byte, error) {
	return e.json.Marshal(&resp)
}
