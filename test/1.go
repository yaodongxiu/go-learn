package main

import "fmt"

func main() {
	var a int
	var e int
	a = 1
	b := a + 1
	b++
	c, d := b, b
	c++
	e = c
	fmt.Println(c, d, e)
	var (
		j struct {
			age int
		}
	)
	j.age = a
}
