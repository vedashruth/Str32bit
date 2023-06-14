package main

import (
	function "Str32bit/functions"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	randomString, err := function.GenerateRandomString()
	if err != nil {
		fmt.Println("Error generating random string:", err)
		return
	}

	fmt.Println("Random String:", randomString)
}
