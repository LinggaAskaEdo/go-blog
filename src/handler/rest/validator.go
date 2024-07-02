package rest

import (
	"errors"

	x "github.com/linggaaskaedo/go-blog/stdlib/error"
)

func ValidateCreateDivision(data *DivisionDataPayload) error {
	if len(data.Name) <= 5 {
		x.Wrap("validateCreateDivision", errors.New("Invalid name length"))
	}

	return nil
}
