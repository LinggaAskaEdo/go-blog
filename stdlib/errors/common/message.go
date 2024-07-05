package common

import "github.com/linggaaskaedo/go-blog/stdlib/errors/entity"

var ErrorMessages = entity.ErrorMessage{
	CodeHTTPBadRequest:          entity.ErrMsgBadRequest,
	CodeHTTPNotFound:            entity.ErrMsgNotFound,
	CodeHTTPUnauthorized:        entity.ErrMsgUnauthorized,
	CodeHTTPInternalServerError: entity.ErrMsgISE,
	CodeHTTPUnmarshal:           entity.ErrMsgBadRequest,
	CodeHTTPMarshal:             entity.ErrMsgISE,
	CodeHTTPConflict:            entity.ErrMsgConflict,
	CodeHTTPForbidden:           entity.ErrMsgForbidden,
	CodeHTTPUnprocessableEntity: entity.ErrMsgUnprocessable,
	CodeHTTPTooManyRequest:      entity.ErrMsgTooManyRequest,
	CodeHTTPServiceUnavailable:  entity.ErrMsgServiceUnavailable,
	CodeHTTPParamDecode:         entity.ErrMsgBadRequest,
	CodeHTTPErrorOnReadBody:     entity.ErrMsgISE,
}
