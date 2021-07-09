package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 标准格式
	var a int // var 变量名 变量类型

	// 批量格式
	var (
		b string
		c = 1.1 // 编译器推导类型的格式
		d byte
		e bool
		f rune
		g complex128
	)

	// 简短格式
	cc := 2.2
	// 注意：由于使用了:=，而不是赋值的=，因此推导声明写法的左值变量必须是没有定义过的变量。
	// 若定义过，将会发生编译错误。
	h, c := 1, 3.3 // 编译器推导类型的格式
	i, c := 2, 4.4
	fmt.Println(a, b, c, cc, d, e, f, g, h, i)

	// output:
	// 0  4.4 2.2 0 false 0 (0+0i) 1 2

	// '1' 类型可以自动在rune和byte间转换适配
	var ii int
	ii = 101
	str := strconv.Itoa(ii)
	for i, v := range str {
		fmt.Println(v)
		fmt.Println(v == '1')
		fmt.Println(str[i] == '1')
	}
}
