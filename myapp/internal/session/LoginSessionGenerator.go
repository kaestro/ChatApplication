// /myapp/internal/session/LoginSessionGenerator.go - Login session generator
package session

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomSessionKey() (string, error) {
	key := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	// base64 encoding to make it a readable string
	return base64.StdEncoding.EncodeToString(key), nil
}
