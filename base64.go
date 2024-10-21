package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Rangga Dwi Mahendra"

	// encode 
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	// decode
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}