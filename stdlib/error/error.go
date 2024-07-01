package error

import "github.com/pkg/errors"

func Wrap(message string, err error) error {
	return errors.Wrap(err, message)
}
