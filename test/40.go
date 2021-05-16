package main

import "fmt"

func main() {
	var arr [3][2]int // 注意该数组中一维有3个元素，二维各有2个元素
	arr[0] = [2]int{1, 2}
	arr[1] = [2]int{3, 4}
	arr[2] = [2]int{5, 6}
	fmt.Println(arr)
	// output: [[1 2] [3 4] [5 6]]
}
