package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// EncryptPwd 加密密码
func EncryptPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

// VerifyPwd 校验密码
func VerifyPwd(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)

	return err == nil
}
