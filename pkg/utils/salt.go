package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func SaltGen() string {
	tmp := make([]byte, 8)
	_, err := rand.Read(tmp)
	if err != nil {
		return "1111111111"
	}
	salt := base64.StdEncoding.EncodeToString(tmp)[:10]
	return salt
}

func Encrypt(origin string) string {
	// 返回长度为 64 的字符串
	hashed := sha256.Sum256([]byte(origin))
	ciphertext := hex.EncodeToString(hashed[:])
	return ciphertext
}
