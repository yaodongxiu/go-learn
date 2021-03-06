package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//done := make(chan int)
	//go func() {
	//	fmt.Println("C语言中文网")
	//	<-done
	//}()
	//done <- 1

	//done := make(chan int, 2)
	//go func() {
	//	fmt.Println("C语言中文网")
	//	done <- 1
	//}()
	//<-done

	//done := make(chan int, 10) // 带10个缓存
	//// 开N个后台打印线程
	//for i := 0; i < cap(done); i++ {
	//	go func() {
	//		fmt.Println("C语言中文网")
	//		done <- 1
	//	}()
	//}
	//// 等待N个后台线程完成
	//for i := 0; i < cap(done); i++ {
	//	<-done
	//}

	runtime.GOMAXPROCS(10)
	var wg sync.WaitGroup
	// 开N个后台打印线程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second)
			fmt.Println("C语言中文网")
			wg.Done()
		}()
	}
	// 等待N个后台线程完成
	wg.Wait()
}
