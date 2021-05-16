package main

import (
	"fmt"
)

func main() {
	done := make(chan int, 1)
	go func() {
		//time.Sleep(2 * time.Second)
		fmt.Println("heartbeat...")
		done <- 1
		//fmt.Println("heartbeat...")
		//<-done
	}()
	fmt.Println("wait")
	<-done
	//done <- 1
	fmt.Println("shutdown")
}
