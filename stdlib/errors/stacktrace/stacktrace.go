package stacktrace

import (
	"fmt"
	"math"
	"runtime"
	"strings"

	"github.com/linggaaskaedo/go-blog/stdlib/errors/stacktrace/cleanpath"
)

type ErrorCode uint32

const NoCode ErrorCode = math.MaxUint32

var CleanPath = cleanpath.RemoveGoPath

type stacktrace struct {
	message  string
	cause    error
	code     ErrorCode
	file     string
	function string
	line     int
}

func NewErrorWithCode(code ErrorCode, msg string, vals ...interface{}) error {
	return create(nil, code, msg, vals...)
}

func Propagate(cause error, msg string, vals ...interface{}) error {
	if cause == nil {
		panic("Propagate called with nil error")
	}

	return create(cause, NoCode, msg, vals...)
}

func PropagateWithCode(cause error, code ErrorCode, msg string, vals ...interface{}) error {
	if cause == nil {
		panic("PropagateWithCode called with nil error")
	}

	return create(cause, code, msg, vals...)
}

func create(cause error, code ErrorCode, msg string, vals ...interface{}) error {
	// If no error code specified, inherit error code from the cause.
	if code == NoCode {
		code = GetCode(cause)
	}

	err := &stacktrace{
		message: fmt.Sprintf(msg, vals...),
		cause:   cause,
		code:    code,
	}

	// Caller of create is NewError or Propagate, so user's code is 2 up.
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return err
	}

	if CleanPath != nil {
		file = CleanPath(file)
	}

	err.file, err.line = file, line

	f := runtime.FuncForPC(pc)
	if f == nil {
		return err
	}

	err.function = shortFuncName(f)

	return err
}

func GetCode(err error) ErrorCode {
	if err, ok := err.(*stacktrace); ok {
		return err.code
	}

	return NoCode
}

func shortFuncName(f *runtime.Func) string {
	// f.Name() is like one of these:
	// - "github.com/palantir/shield/package.FuncName"
	// - "github.com/palantir/shield/package.Receiver.MethodName"
	// - "github.com/palantir/shield/package.(*PtrReceiver).MethodName"
	longName := f.Name()

	withoutPath := longName[strings.LastIndex(longName, "/")+1:]
	withoutPackage := withoutPath[strings.Index(withoutPath, ".")+1:]

	shortName := withoutPackage
	shortName = strings.Replace(shortName, "(", "", 1)
	shortName = strings.Replace(shortName, "*", "", 1)
	shortName = strings.Replace(shortName, ")", "", 1)

	return shortName
}

func (st *stacktrace) Error() string {
	type _stacktrace *stacktrace

	return fmt.Sprintf("%+v", _stacktrace(st))
}
