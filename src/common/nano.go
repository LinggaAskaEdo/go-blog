package common

import (
	"strings"

	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/pkg/errors"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyz"
	length   = 13
)

func NewPID() (string, error) {
	return nanoid.Generate(alphabet, length)
}

func MustPID() string {
	return nanoid.MustGenerate(alphabet, length)
}

func ValidatePID(fieldName, id string) error {
	if id == "" {
		return errors.Errorf("%s cannot be blank", fieldName)
	}

	if len(id) != length {
		return errors.Errorf("%s should be %d characters long", fieldName, length)
	}

	if strings.Trim(id, alphabet) != "" {
		return errors.Errorf("%s has invalid characters", fieldName)
	}

	return nil
}
