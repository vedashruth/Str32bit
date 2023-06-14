package functions

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomString() (string, error) {
	randomBytes := make([]byte, 12)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	fmt.Println(randomString)
	randomString = randomString[:12]
	fmt.Println(randomString)

	return randomString, nil
}
