package main

import "fmt"

func main() {
	const str = `第一行
第二行
第三行
\r\n
`
	a := "123"
	fmt.Println(str + a)
}
