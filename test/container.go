package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"sort"
	"strings"
	"sync"
)

func main() {
	tracer := "死神来了, 狙击bye bye"
	comma := strings.Index(tracer, ", ")
	pos := strings.Index(tracer[comma:], "死神")

	fmt.Println(comma, pos, tracer[comma+pos:])
	fmt.Println(tracer[14:])
	fmt.Println(tracer[0:7])

	tracer2 := "123456789"
	fmt.Println(tracer2[1:5])
	fmt.Println(tracer2[:5])
	fmt.Println(tracer2[0:5])
	fmt.Println(tracer2[0:])
	fmt.Println(tracer2[:])

	a := make([]int, 2, 6)
	b := append(a, 1)
	fmt.Println(a, b)
	a = append(a, []int{2, 3, 4}...)
	fmt.Println(a)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	fmt.Println(999999)
	copy(slice1[2:], slice2)
	fmt.Println(slice1)

	slice1 = []int{1, 2, 3, 4, 5}
	slice2 = []int{5, 4, 3}
	copy(slice2, slice1)
	fmt.Println(slice2)

	// 设置元素数量为1000
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
	fmt.Println()

	a = []int{1, 2, 3}
	a = a[:copy(a, a[1:])] // 删除开头1个元素
	fmt.Println(a)

	map1 := make(map[string]map[string]int)
	map2 := make(map[string]int)
	map2["bbb"] = 222
	map1["aaa"] = map2
	fmt.Println(map1)

	for _, v := range map2 {
		fmt.Println(v)
	}

	testMap := make(map[string]int)
	// 准备map数据
	testMap["route"] = 66
	testMap["brazil"] = 4
	testMap["china"] = 960
	// 声明一个切片保存map数据
	var testMapList []string
	// 将map数据遍历复制到切片中
	for k := range testMap {
		testMapList = append(testMapList, k)
	}
	// 对切片进行排序
	sort.Strings(testMapList)
	// 输出
	fmt.Println(testMapList)

	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	fmt.Println(scene.Load("london"))
	scene.Range(func(key interface{}, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	testList := list.New()
	testList.PushFront("a")
	element := testList.PushBack(666)
	testList.InsertAfter("end", element)
	testList.InsertBefore("front", element)
	testList.Remove(element)
	for i := testList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	r := ring.New(10) //初始长度10
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	for i := 0; i < r.Len(); i++ {
		fmt.Println(r.Value)
		r = r.Next()
	}
	r = r.Move(6)
	fmt.Println(r.Value) //6
	r1 := r.Unlink(19)   //移除19%10=9个元素
	for i := 0; i < r1.Len(); i++ {
		fmt.Println(r1.Value)
		r1 = r1.Next()
	}
	fmt.Println(r.Len())  //10-9=1
	fmt.Println(r1.Len()) //9
}
