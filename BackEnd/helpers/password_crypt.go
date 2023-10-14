package helpers

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) ([]byte, error) {
	passwordInBytes := []byte(password)
	encryptedPassword, err := bcrypt.GenerateFromPassword(passwordInBytes, 10)
	return encryptedPassword, err
}

func PasswordMatchesTheHash(password string, hash string) bool {
	passwordInBytes := []byte(password)
	hashInBytes := []byte(hash)
	result := bcrypt.CompareHashAndPassword(hashInBytes, passwordInBytes)
	return result == nil
}
