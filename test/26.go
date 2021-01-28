package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Println(221312)
		log.Fatal(err)
	}

	// Output:
	// some io.Reader stream to be read

}
