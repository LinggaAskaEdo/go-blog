package entity

import "github.com/linggaaskaedo/go-blog/stdlib/errors/stacktrace"

var (
	ErrCode      = stacktrace.GetCode
	NewWithCode = stacktrace.NewErrorWithCode
	Wrap         = stacktrace.Propagate
	WrapWithCode = stacktrace.PropagateWithCode
)

type (
	Code         = stacktrace.ErrorCode
	ErrorMessage map[Code]Message

	Message struct {
		StatusCode    int    `json:"status_code"`
		EN            string `json:"en"`
		ID            string `json:"id"`
		HasAnnotation bool
	}
)
