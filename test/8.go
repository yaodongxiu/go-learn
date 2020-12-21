package main

import (
	"encoding/base64"
	"fmt"
)

// encode and decode
func main() {
	message := "sdfsdfdsf"
	encodeMessage := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(encodeMessage)
	data, err := base64.StdEncoding.DecodeString(encodeMessage)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
