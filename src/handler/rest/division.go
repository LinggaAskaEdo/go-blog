package rest

import (
	"errors"
	"io"
	"net/http"

	"github.com/linggaaskaedo/go-blog/src/business/dto"
	"github.com/linggaaskaedo/go-blog/src/common"
)

func (e *rest) CreateDivision(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		e.httpRespError(w, r, http.StatusBadRequest, err)
		return
	}

	var requestBody DivisionCreateRequest
	if err = e.json.Unmarshal(body, &requestBody); err != nil {
		e.httpRespError(w, r, http.StatusBadRequest, err)
		return
	}

	if requestBody.Data.Division == nil {
		e.httpRespError(w, r, http.StatusBadRequest, errors.New("Invalid payload"))
		return
	}

	data := *&requestBody.Data.Division

	validationErr := ValidateCreateDivision(data)
	if validationErr != nil {
		e.httpRespError(w, r, http.StatusBadRequest, validationErr)
		return
	}

	divisionDTO := dto.DivisionDTO{
		PublicID: common.MustPID(),
		Name:     data.Name,
	}

	result, err := e.uc.Division.CreateDivision(ctx, divisionDTO)
	if err != nil {
		e.httpRespError(w, r, http.StatusInternalServerError, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusCreated, result, nil)
}
