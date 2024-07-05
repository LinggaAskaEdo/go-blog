package rest

import (
	"database/sql"
	"io"
	"net/http"
	"time"

	"github.com/linggaaskaedo/go-blog/src/business/entity"
	"github.com/linggaaskaedo/go-blog/src/common"
	commonerr "github.com/linggaaskaedo/go-blog/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-blog/stdlib/errors/entity"
)

func (e *rest) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		e.httpRespError(w, r, x.WrapWithCode(err, commonerr.CodeHTTPUnprocessableEntity, "body_err"))
		return
	}

	var requestBody UserCreateRequest
	if err = e.json.Unmarshal(body, &requestBody); err != nil {
		e.httpRespError(w, r, x.WrapWithCode(err, commonerr.CodeHTTPUnmarshal, "unmarshall_err"))
		return
	}

	if requestBody.Data.User == nil {
		e.httpRespError(w, r, x.NewWithCode(commonerr.CodeHTTPBadRequest, "invalid_payload"))
		return
	}

	data := requestBody.Data.User

	validationErr := data.Validate()
	if validationErr != nil {
		e.httpRespError(w, r, x.WrapWithCode(validationErr, commonerr.CodeHTTPBadRequest, "validation_error"))
		return
	}

	divisionID, err := common.MixerDecode(data.DivisionID)
	if err != nil {
		e.httpRespError(w, r, x.WrapWithCode(err, commonerr.CodeHTTPNotFound, "decode_error"))
		return
	}

	hashPassword, err := common.HashPassword(data.Password)
	if err != nil {
		e.httpRespError(w, r, x.WrapWithCode(err, commonerr.CodeHTTPBadRequest, "hash_password"))
		return
	}

	_, err = e.uc.Division.GetDivisioByID(ctx, divisionID)
	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	userEntity := entity.User{
		Username: data.Username,
		Email:    data.Email,
		Phone:    data.Phone,
		Division: entity.Division{
			ID: divisionID,
		},
		Password:  hashPassword,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	result, err := e.uc.User.CreateUser(ctx, userEntity)
	if err != nil {
		e.httpRespError(w, r, err)
		return
	}

	e.httpRespSuccess(w, r, http.StatusCreated, result, nil)
}

// func (e *rest) GetUserByID(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	userID := params["userID"]
// 	e.log.Debug().Str("user_param", userID).Send()

// 	vid, err := strconv.ParseInt(userID, 10, 64)
// 	if err != nil {
// 		e.httpRespError(w, r, http.StatusBadRequest, err)
// 		return
// 	}

// 	var cacheControl dto.CacheControl
// 	if r.Header.Get(preference.CacheControl) == preference.CacheMustRevalidate {
// 		cacheControl.MustRevalidate = true
// 	} else if r.Header.Get(preference.CacheControl) == "must-db-revalidate" {
// 		cacheControl.MustDbValidate = true
// 	}

// 	result, err := e.uc.User.GetUserByUserID(r.Context(), cacheControl, vid)
// 	if err != nil {
// 		e.httpRespError(w, r, http.StatusInternalServerError, err)
// 		return
// 	}

// 	e.httpRespSuccess(w, r, http.StatusOK, result, nil)
// }
