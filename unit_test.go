package str32bit

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	_, err := GenerateRandomString()
	if err != nil {
		fmt.Println("Error generating random string:", err)
		return
	}
}
