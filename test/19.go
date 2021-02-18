package main

import "fmt"

func main() {
	var a interface{}
	a = 100
	// 类型断言（Type Assertion）
	value, ok := a.(int)
	fmt.Println(value, ok)

	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
}
