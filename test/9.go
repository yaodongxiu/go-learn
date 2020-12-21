package main

import (
	"fmt"
	"strings"
)

func main() {
	tracer := "死神来了, 狙击bye bye"
	comma := strings.Index(tracer, ", ")
	pos := strings.Index(tracer[comma:], "死神")

	fmt.Println(comma, pos, tracer[comma+pos:])
	fmt.Println(tracer[14:])
	fmt.Println(tracer[0:7])
}
