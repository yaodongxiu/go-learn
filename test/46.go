package main

import (
	"fmt"
	"time"
)

func main() {
	cc := make(chan int, 1)
	i := 0
	for {
		select {
		case v := <-cc:
			fmt.Println(v)
		case <-time.After(2 * time.Second):
			fmt.Println("haha")
			i++
			cc <- i
		}
	}
}
