package error

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/linggaaskaedo/go-blog/stdlib/errors/common"
	"github.com/linggaaskaedo/go-blog/stdlib/errors/entity"
	preference "github.com/linggaaskaedo/go-blog/stdlib/preference"
)

var svcError map[ServiceType]entity.ErrorMessage

type ServiceType int

const (
	COMMON ServiceType = 1
)

type AppError struct {
	Code       entity.Code `json:"code"`
	Message    string      `json:"message"`
	DebugError *string     `json:"debug,omitempty"`
	sys        error
}

func init() {
	svcError = map[ServiceType]entity.ErrorMessage{
		COMMON: common.ErrorMessages,
	}
}

func Compile(service ServiceType, err error, lang string, debugMode bool) (int, AppError) {
	var debugErr *string

	if debugMode {
		errStr := err.Error()
		if len(errStr) > 0 {
			debugErr = &errStr
		}
	}

	code := entity.ErrCode(err)

	if errMessage, ok := svcError[COMMON][code]; ok {
		msg := errMessage.ID
		if lang == preference.LangEN {
			msg = errMessage.EN
		}

		return errMessage.StatusCode, AppError{
			Code:       code,
			Message:    msg,
			sys:        err,
			DebugError: debugErr,
		}
	}

	if errMessages, ok := svcError[service]; ok {
		if errMessage, ok := errMessages[code]; ok {
			msg := errMessage.ID
			if lang == preference.LangEN {
				msg = errMessage.EN
			}

			if errMessage.HasAnnotation {
				args := fmt.Sprintf("%q", err.Error())
				if start, end := strings.LastIndex(args, `{{`), strings.LastIndex(args, `}}`); start > -1 && end > -1 {
					args = strings.TrimSpace(args[start+2 : end])
					msg = fmt.Sprintf(msg, args)
				} else {
					// Deprecated: Old style still support to handle broken code
					index := strings.Index(args, `\n`)
					if index > 0 {
						args = strings.TrimSpace(args[1:index])
					}
					msg = fmt.Sprintf(msg, args)
				}
			}

			if code == common.CodeHTTPValidatorError {
				if err.Error() != "" {
					msg = strings.Split(err.Error(), "\n ---")[0]
				}
			}

			return errMessage.StatusCode, AppError{
				Code:       code,
				Message:    msg,
				sys:        err,
				DebugError: debugErr,
			}
		}

		return http.StatusInternalServerError, AppError{
			Code:       code,
			Message:    "error message not defined!",
			sys:        err,
			DebugError: debugErr,
		}
	}

	return http.StatusInternalServerError, AppError{
		Code:       code,
		Message:    "service error not defined!",
		sys:        err,
		DebugError: debugErr,
	}
}

// var (
// 	CleanPath = cleanpath.RemoveGoPath
// )

// type stacktrace struct {
// 	message  string
// 	cause    error
// 	code     int
// 	file     string
// 	line     int
// 	function string
// }

// type AppError struct {
// 	// Code       entity.Code `json:"code"`
// 	Message    string  `json:"message"`
// 	DebugError *string `json:"debug,omitempty"`
// 	sys        error
// }

// func (st *stacktrace) Error() string {
// 	type _stacktrace *stacktrace

// 	return fmt.Sprintf("%+v", _stacktrace(st))
// }

// func Wrap(message string, err error) error {
// 	return errors.Wrap(err, message)
// }

// func WrapWithCode(msg string, cause error, code int, vals ...interface{}) error {
// 	return create(cause, code, msg, vals...)

// 	// return &stacktrace{
// 	// 	message: message,
// 	// 	err:     err,
// 	// 	code:    code,
// 	// }
// }

// func create(cause error, code int, msg string, vals ...interface{}) error {
// 	// If no error code specified, inherit error code from the cause.
// 	// if code == 0 {
// 	// 	code = GetCode(cause)
// 	// }

// 	code = GetCode(cause)

// 	err := &stacktrace{
// 		message: fmt.Sprintf(msg, vals...),
// 		cause:   cause,
// 		code:    code,
// 	}

// 	// Caller of create is NewError or Propagate, so user's code is 2 up.
// 	pc, file, line, ok := runtime.Caller(2)
// 	if !ok {
// 		return err
// 	}

// 	if CleanPath != nil {
// 		file = CleanPath(file)
// 	}

// 	err.file, err.line = file, line

// 	f := runtime.FuncForPC(pc)
// 	if f == nil {
// 		return err
// 	}

// 	err.function = shortFuncName(f)

// 	return err
// }

// func GetCode(err error) int {
// 	if err, ok := err.(*stacktrace); ok {
// 		return err.code
// 	}

// 	return 0
// }

// func shortFuncName(f *runtime.Func) string {
// 	// f.Name() is like one of these:
// 	// - "github.com/palantir/shield/package.FuncName"
// 	// - "github.com/palantir/shield/package.Receiver.MethodName"
// 	// - "github.com/palantir/shield/package.(*PtrReceiver).MethodName"
// 	longName := f.Name()

// 	withoutPath := longName[strings.LastIndex(longName, "/")+1:]
// 	withoutPackage := withoutPath[strings.Index(withoutPath, ".")+1:]

// 	shortName := withoutPackage
// 	shortName = strings.Replace(shortName, "(", "", 1)
// 	shortName = strings.Replace(shortName, "*", "", 1)
// 	shortName = strings.Replace(shortName, ")", "", 1)

// 	return shortName
// }

// func Compile(err error) (int, AppError) {
// 	var debugErr *string

// 	// Get Error Code
// 	code := entity.ErrCode(err)

// 	// Get Common Error
// 	if errMessage, ok := svcError[COMMON][code]; ok {
// 		msg := errMessage.ID
// 		if lang == header.LangEN {
// 			msg = errMessage.EN
// 		}
// 		return errMessage.StatusCode, AppError{
// 			Code:       code,
// 			Message:    msg,
// 			sys:        err,
// 			DebugError: debugErr,
// 		}
// 	}
// }
