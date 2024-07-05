package common

import "github.com/linggaaskaedo/go-blog/stdlib/errors/entity"

const (
	// Code HTTP Handler
	CodeHTTPBadRequest = entity.Code(iota + 800)
	CodeHTTPNotFound
	CodeHTTPUnauthorized
	CodeHTTPInternalServerError
	CodeHTTPUnmarshal
	CodeHTTPMarshal
	CodeHTTPConflict
	CodeHTTPForbidden
	CodeHTTPUnprocessableEntity
	CodeHTTPTooManyRequest
	CodeHTTPValidatorError
	CodeHTTPServiceUnavailable
	CodeHTTPParamDecode
	CodeHTTPErrorOnReadBody
)
