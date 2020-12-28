package main

import (
	"fmt"
	"math/rand"
	"time"
)

func produce(header string, channel chan<- string) {
	for {
		channel <- fmt.Sprintf("%s: %d", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func consume(channel <-chan string) {
	for {
		fmt.Println(<-channel)
	}
}

func main() {
	channel := make(chan string)
	go produce("dooog", channel)
	go produce("caaat", channel)
	consume(channel)
}
