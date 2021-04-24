package main

import (
	"fmt"
)

type Stu struct {
	AgeNumber int
	Name      string `json:"name"`
}

// []Stu和[]*Stu 对于在传递切片的时候，修改切片元组的值没有区别，唯一的区别是指针切片[]*Stu的数据都产生了逃逸，结构体切片[]Stu的数据在栈里
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
