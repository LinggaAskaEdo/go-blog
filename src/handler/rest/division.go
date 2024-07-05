package rest

import (
	"database/sql"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/linggaaskaedo/go-blog/src/business/entity"
	"github.com/linggaaskaedo/go-blog/src/common"
	commonerr "github.com/linggaaskaedo/go-blog/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-blog/stdlib/errors/entity"
)

func (e *rest) CreateDivision(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		e.httpRespError(w, r, x.WrapWithCode(err, commonerr.CodeHTTPUnprocessableEntity, "body_err"))
		return
	}

	var requestBody DivisionCreateRequest
	if err = e.json.Unmarshal(body, &requestBody); err != nil {
		e.httpRespError(w, r, x.WrapWithCode(err, commonerr.CodeHTTPUnmarshal, "unmarshall_err"))
		return
	}

	if requestBody.Data.Division == nil {
		e.httpRespError(w, r, x.NewWithCode(commonerr.CodeHTTPBadRequest, "invalid_payload"))
		return
	}

	data := requestBody.Data.Division

	validationErr := data.Validate()
	if validationErr != nil {
		e.httpRespError(w, r, x.WrapWithCode(validationErr, commonerr.CodeHTTPBadRequest, "validation_error"))
		return
	}

	divisionEntity := entity.Division{
		Name:      data.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	result, err := e.uc.Division.CreateDivision(ctx, divisionEntity)
	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusCreated, result, nil)
}

func (e *rest) GetDivisioByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	params := mux.Vars(r)
	varID := params["divisionID"]

	divisionID, err := common.MixerDecode(varID)
	if err != nil {
		e.httpRespError(w, r, x.Wrap(err, "decode_division_id"))
		return
	}

	result, err := e.uc.Division.GetDivisioByID(ctx, divisionID)
	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusOK, result, nil)
}
