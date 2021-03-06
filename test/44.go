package main

import (
	"fmt"
)

type Stu struct {
	AgeNumber int
	Name      string `json:"name"`
}

// []Stu和[]*Stu 对于在传递切片的时候，修改切片元组的值没有区别，唯一的区别是指针切片[]*Stu的数据都产生了逃逸，结构体切片[]Stu的数据在栈里
// 除了指针，只有slice map chan是引用传递类型, 结构体不是引用传递类型，所以在将结构体实例作为函数参数时会复制一份结构体实例
func main() {
	students := make([]Stu, 0, 3)
	for i := 0; i < 10; i++ {
		students = append(students, Stu{AgeNumber: i})
	}
	fmt.Println(students[0])
	ttt(students[0])
	fmt.Println(students[0])
	ttt2(students)
	fmt.Println(students[0])

	students34 := make([]*Stu, 0, 3)
	for i := 0; i < 10; i++ {
		students34 = append(students34, &Stu{AgeNumber: i})
	}
	fmt.Println(students34[0])
	ttt3(students34[0])
	fmt.Println(students34[0])
	ttt4(students34)
	fmt.Println(students34[0])

	stu111 := Stu{25, "yaodx"}
	stu222 := stu111
	stu222.AgeNumber = 30
	fmt.Println(&stu111, &stu222)
	stu333 := &Stu{18, "mm"}
	stu444 := *stu333
	stu444.AgeNumber = 30
	fmt.Println(stu444)
	fmt.Println(stu333, &stu444)

	aaaaa := [2]int{1, 2}
	ttt5(aaaaa)
	fmt.Println(aaaaa)
}

func ttt(s Stu) {
	s.AgeNumber = 111
}

func ttt2(stu []Stu) {
	stu[0].AgeNumber = 111
}

func ttt3(s *Stu) {
	s.AgeNumber = 111
}

func ttt4(stu []*Stu) {
	stu[0].AgeNumber = 111
}

func ttt5(a [2]int) {
	a[0] = 333
}
