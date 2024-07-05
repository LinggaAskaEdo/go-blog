package dto

import apperr "github.com/linggaaskaedo/go-blog/stdlib/errors"

type Meta struct {
	Path       string           `json:"path"`
	StatusCode int              `json:"status_code"`
	Status     string           `json:"status"`
	Message    string           `json:"message"`
	Error      *apperr.AppError `json:"error,omitempty" swaggertype:"primitive,object"`
	Timestamp  string           `json:"timestamp"`
}
