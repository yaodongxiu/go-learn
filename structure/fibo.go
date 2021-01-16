package main

import "fmt"

func fibo(x int) int {
	if x == 1 {
		return 0
	}
	if x == 2 {
		return 1
	}
	return fibo(x-1) + fibo(x-2)
}

func main() {
	fmt.Println(fibo(9))
}
