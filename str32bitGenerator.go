package str32bit

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString() (string, error) {
	randomBytes := make([]byte, 12)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	randomString = randomString[:12]
	return randomString, nil
}
