package main

import "fmt"

var step int

func move(n int, src, dst string) {
	step++
	fmt.Println(fmt.Sprintf("第%d步 块[%d] %s => %s", step, n, src, dst))
}

func hanoi(n int, from, by, to string) {
	if n <= 1 {
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
