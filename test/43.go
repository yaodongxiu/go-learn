package main

import (
	"fmt"
)

func test1() {

	//defer func() {
	//	defer func() {
	//		if err := recover(); err != nil {
	//			fmt.Println(err)
	//		}
	//	}()
	//	panic("test panic 222")
	//}()
	//
	//panic("test panic 111")

	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	}
	//}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		// 有 panic 也有 recover, 恢复了相当于啥都没有
	}()

	panic("panic 000") // 从宕机点退出当前函数后，所以下面一行无法执行
	fmt.Println(555)   //
}

func main() {
	test1() // recover后从宕机点退出当前函数后继续执行主函数的代码
	println(666)
}
