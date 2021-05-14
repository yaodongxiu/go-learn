package main

import "fmt"

type aaaa struct {
	ID uint
}

func main() {
	//a := &aaaa{}
	//fmt.Println(a)
	//g := a
	//fmt.Println(a, g)
	//var b *aaaa
	//fmt.Println(b)
	//a = nil
	//fmt.Println(a)
	//c := new(aaaa)
	//fmt.Println(c)

	d := []int{}
	fmt.Println(d)
	fmt.Println(d == nil)
	e := make([]int, 0)
	fmt.Println(e)
	fmt.Println(e == nil)
	var f []int
	fmt.Println(f)
	fmt.Println(f == nil)
	f = nil
	fmt.Println(f)
	asd := []int{111}
	kkk(asd)
	fmt.Println(asd)
	bsd := map[int]int{0: 111}
	ppp(bsd)
	fmt.Println(bsd)
}

func kkk(a []int) {
	a[0] = 999
	a = append(a, 888)
}

func ppp(b map[int]int) {
	b[0] = 555
	b[1] = 888
}
