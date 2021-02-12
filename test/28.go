package main

import "fmt"

func main() {
	var a []int
	a = append(a, 1)                 // 追加1个元素
	a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	fmt.Println(a)

	// 从开头位置删除
	a = []int{1, 2, 3}
	a = a[1:] // 删除开头1个元素
	//a = a[N:] // 删除开头N个元素

	// 从中间位置i删除
	//i:= 10
	//a = []int{1, 2, 3, ...}
	//a = append(a[:i], a[i+1:]...) // 删除中间1个元素
	//a = append(a[:i], a[i+N:]...) // 删除中间N个元素
	//a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	//a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素

	// 从尾部删除
	a = a[:len(a)-1] // 删除尾部1个元素
	//a = a[:len(a)-N] // 删除尾部N个元素
}
