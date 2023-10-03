package utils

import "golang.org/x/crypto/bcrypt"

func Hash(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), 15)
	return string(bytes), err
}

func CompareHash(plain, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
	return err == nil
}
