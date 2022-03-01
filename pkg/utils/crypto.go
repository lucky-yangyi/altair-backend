package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

func GenerateMd5(input string) (output string) {
	hash := md5.New()
	hash.Write([]byte(input))
	output = fmt.Sprintf("%x", hash.Sum(nil))
	return
}

func Hmac256AndBase256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	sha := hex.EncodeToString(h.Sum(nil))

	log.Println("sha256 ----->", sha)

	return base64.StdEncoding.EncodeToString([]byte(sha))
}

//Sha256加密
func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
