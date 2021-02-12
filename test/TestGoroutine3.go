package main

import "fmt"

func main() {
	channel := make(chan int, 3)
	fmt.Println(len(channel))
	channel <- 1
	channel <- 2
	channel <- 3
	fmt.Println(len(channel))
}
