package utils

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) (string, error) {
	return HashAndSalt([]byte(password))
}
func CheckPassword(password, hash string) bool {
	return ComparePasswords(hash, []byte(password))
}
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
func HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
