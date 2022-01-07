package security

import (
	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func VerifyPassword(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

func GenerateUUID() string {

	uuidWithHyphen := uuid.New()
	// uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

	return uuidWithHyphen.String()
}
