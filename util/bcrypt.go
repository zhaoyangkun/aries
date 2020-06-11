package util

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// 加密密码
func EncryptPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hash), err
}

// 校验密码
func VerifyPwd(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		return false
	}
	return true
}
