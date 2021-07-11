package main

import "fmt"

// 切片使用注意事项
func main() {
	fmt.Println("注意在使用 append(a[:i], a[i+j:]...)的时候，也会对切片a进行了修改，对切片a进行了污染")
	asddf := []int{1, 2, 3, 4}
	i, j := 1, 2
	b := append(asddf[:i], asddf[i+j:]...)
	fmt.Println("b = ", b)
	asddf[0] = 222
	fmt.Println("b = ", b)
	fmt.Println(asddf)
	asddf = asddf[:len(asddf)-j]
	fmt.Println(asddf)
	fmt.Println("正确用法需要做如下操作")
	asddf = []int{1, 2, 3, 4}
	ttttmp := make([]int, i, i)
	copy(ttttmp, asddf[:i])
	ttttmp = append(ttttmp, asddf[i+1:]...)
	fmt.Println(asddf)

	// 同理，把切片加入切片数组也一样
	var res [][]int
	target := []int{1, 2, 3}
	fmt.Println("错误做法如下")
	for i := 0; i < 3; i++ {
		target[2] = target[2] + i
		res = append(res, target)
	}
	fmt.Println(res)
	fmt.Println("正确做法如下")
	res = nil
	for i := 0; i < 3; i++ {
		target[2] = target[2] + i
		tmp := make([]int, len(target), len(target))
		copy(tmp, target)
		res = append(res, tmp)
	}
	fmt.Println(res)
}
