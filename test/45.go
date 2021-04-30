package main

import "fmt"

type aaaa struct {
	ID uint
}

func main() {
	a := &aaaa{}
	fmt.Println(a)
	g := a
	fmt.Println(a, g)
	var b *aaaa
	fmt.Println(b)
	a = nil
	fmt.Println(a)
	c := new(aaaa)
	fmt.Println(c)

	d := []int{}
	fmt.Println(d)
	e := make([]int, 0)
	fmt.Println(e)
	fmt.Println(e == nil)
	var f []int
	fmt.Println(f)
	fmt.Println(f == nil)
	f = nil
	fmt.Println(f)
}
