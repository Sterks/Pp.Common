package common

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
)

//GetHash Получить sha256 хеш .
func GetHash(file io.ReadCloser) string {
	// var g io.Writer
	var hasher = sha256.New()
	// repeatStream := io.TeeReader(file, hasher)
	// _, err := io.Copy(hasher, file)
	// _, err := hasher.Write(file)
	_, err := io.Copy(hasher, file)
	if err != nil {
		log.Printf("Не удалось расчитать hash - %v", err)
	}
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
