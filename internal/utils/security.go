package utils

import "golang.org/x/crypto/bcrypt"

// Hash Password
func HashPassword(password string) (string, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(newPassword), nil
}

// Check passoword
func VerifyPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
