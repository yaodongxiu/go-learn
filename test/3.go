package main

import "fmt"

func GetData() (int, int, int) {
	return 100, 200, 300
}

func main() {
	a, _, _ := GetData()
	_, b, _ := GetData()
	_, _, c := GetData()
	fmt.Println(a, b, c)
}
