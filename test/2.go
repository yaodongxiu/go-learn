package main

import "fmt"

func main() {
	a := 100
	b := 200
	//a = a ^ b
	//b = b ^ a
	//a = a ^ b
	a, b = b, a
	fmt.Println(a, b)

	c := -8
	c = c << 2
	fmt.Println(c)
	c = c >> 6
	fmt.Println(c)
}
