package str32bit

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	randomString, err := GenerateRandomString()
	if err != nil {
		fmt.Println("Error generating random string:", err)
		return
	}

	fmt.Println("Random String:", randomString)
}
