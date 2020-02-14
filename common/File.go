package common

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
)

//GetHash Получить md5 хеш .
func GetHash(file io.Reader) (string, io.Writer) {
	var g io.Writer
	var hasher = sha256.New()
	// repeatStream := io.TeeReader(file, hasher)
	_, err := io.Copy(g, file)
	if err != nil {
		log.Printf("Не удалось расчитать hash - %v", err)
	}
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash, g
}
