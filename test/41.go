package main

import "fmt"

func main() {
	s := make([]int, 12, 64)
	fmt.Println(s, len(s), cap(s))
	c := make(chan int, 64)
	c <- 1
	fmt.Println(len(c), cap(c))
	m := make(map[string]string, 64)
	fmt.Println(len(m))

	// output:
	// [0 0 0 0 0 0 0 0 0 0 0 0] 12 64
	// 1 64
	// 0
}
