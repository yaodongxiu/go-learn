package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Println(a, b, c)
	goos := runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
