package common

import (
	"golang.org/x/crypto/bcrypt"

	x "github.com/linggaaskaedo/go-blog/stdlib/errors/entity"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return string(bytes), x.Wrap(err, "generate_hash")
	}

	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}

	return true
}
