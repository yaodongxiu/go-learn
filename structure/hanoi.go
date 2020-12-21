package main

import "fmt"

var step int

func move(n int, from string, to string) {
	step++
	fmt.Println(fmt.Sprintf("第%d步 块[%d] %s => %s", step, n, from, to))
}

func hanoi(n int, from string, by string, to string) {
	if n == 1 {
		move(1, from, to)
	} else {
		hanoi(n-1, from, to, by)
		move(n, from, to)
		hanoi(n-1, by, from, to)
	}
}

func main() {
	hanoi(3, "a", "b", "c")
}
