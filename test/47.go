package main

import "fmt"

func main() {
	fmt.Println("注意在使用 append(a[:i], a[i+j:]...)的时候，也会对切片a进行了修改，对切片a进行了污染")
	asddf := []int{1, 2, 3, 4}
	i, j := 1, 2
	_ = append(asddf[:i], asddf[i+j:]...)
	asddf = asddf[:len(asddf)-j]
	fmt.Println(asddf)
	fmt.Println("正确用法需要做如下操作")
	ttttmp := make([]int, i)
	copy(ttttmp, asddf[:i])
	ttttmp = append(ttttmp, asddf[i+1:]...)
	fmt.Println(ttttmp)
}
