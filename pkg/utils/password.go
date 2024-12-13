package utils

import (
	"github.com/jumayevgadam/tsu-toleg/pkg/errlst"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword method.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errlst.NewBadRequestError("error hashing password in [utils][HashPassword]")
	}

	return string(hashed), nil
}

// CheckPassword method checks hashed password with login(time) parol.
func CheckAndComparePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
