package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(passwordHashed, password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHashed),
		[]byte(password),
	)

	return err == nil
}
