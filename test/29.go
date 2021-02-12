package main

import "fmt"

func test() (a, b int) {
	a = 1
	b = 2
	return
}

func main() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
	aaa, bbb := test()
	println(aaa, bbb)
}
