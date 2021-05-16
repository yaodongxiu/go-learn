package main

import "fmt"

// 结论：
// nil slice等于nil 未分配内存，empty slice已分配内存，内容指向一块地址
// nil slice、empty slice都可以使用append添加元素
// len()、cap()对nil slice、empty slice都合法，可能正是因为如此，append才合法
func main() {
	fmt.Print("\nnil slice\n")
	var slice []int
	fmt.Print(len(slice))
	fmt.Print(cap(slice))
	s := append(slice, 1)
	fmt.Print(s, (slice == nil))

	fmt.Print("\n/\n")
	fmt.Print("\nempty slice\n")
	s1 := make([]int, 0, 0)
	s1v := append(s1, 1, 2)
	fmt.Print(s1v)
	fmt.Print(s1v[1])
	//fmt.Print(s1v[len(s1v)])//引发数组越界panic
	var str = "abcd"
	fmt.Print(string(str[1]))
	//fmt.Print(string(str[len(str)])) 引发数组越界panic

	fmt.Print("\n/\n")
	fmt.Print("\nnil map\n")
	var m map[int]int
	fmt.Print(len(m))
	//m1 := map[0] = 0 //报错

	fmt.Print("\n/\n")
	fmt.Print("\nempty map\n")
	m1 := make(map[int]int)
	fmt.Println(len(m1))
	m1[0] = 1
	fmt.Println(m1)
	fmt.Println(len(m1))
	fmt.Println(m1[999]) //不会引发越界错误，为类型的零值
}
