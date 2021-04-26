package main

import "fmt"

// 车轮
type Wheel struct {
	Size int
}

// 引擎
type Engine struct {
	Size  int
	Power int    // 功率
	Type  string // 类型
}

// 车
type Car struct {
	Wheel
	Engine
}

func main() {
	c := Car{
		// 初始化轮子
		Wheel: Wheel{
			Size: 18,
		},
		// 初始化引擎
		Engine: Engine{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)
	c.Engine.Size = 100
	fmt.Printf("%+v\n", c)

	w1 := Wheel{1}
	var w2 Wheel
	w2 = w1 // 这里其实是结构体拷贝，不是同一个结构体的引用
	if &w1 == &w2 {
		fmt.Println("&w1 equals to &w2")
	}

	if w1 == w2 { // 比较俩个结构体的属性
		fmt.Println("w1 equals to w2")
	}

	w2.Size = 2
	if w1 == w2 {
		fmt.Println("w1 equals to w2 after changing w2")
	}

	a := [3]int{1, 2, 3}
	b := a
	b[0] = 666
	fmt.Println(a)
}
