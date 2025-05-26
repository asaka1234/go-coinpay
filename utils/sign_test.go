package utils

import (
	"fmt"
	"testing"
)

func TestSignDeposit(t *testing.T) {

	paramsMap := map[string]interface{}{
		"a":  123,
		"b":  "aaa",
		"aB": "cccc",
	}
	aKey := "12345"

	result := SignDeposit(paramsMap, aKey)

	fmt.Printf("%s\n", result)

}
