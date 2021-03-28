package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var cv = sync.NewCond(&sync.Mutex{})

func takeStep1() {
	fmt.Println(111)
	cv.L.Lock()
	fmt.Println(222)
	cv.Wait()
	fmt.Println(333)
	cv.L.Unlock()
	fmt.Println(444)
}

func takeStep2() {
	fmt.Println(555)
	cv.L.Lock()
	fmt.Println(666)
	cv.Wait()
	fmt.Println(777)
	cv.L.Unlock()
	fmt.Println(888)
}

func main() {
	runtime.GOMAXPROCS(4)
	go takeStep1()
	go takeStep2()
	time.Sleep(time.Second * 3)

	go func() {
		time.Sleep(time.Second * 10)
		cv.L.Unlock()
	}()

	for {
		fmt.Println(0)
		cv.Broadcast() // 唤醒全部协程
		//cv.Signal()    // 唤醒一个协程
		time.Sleep(time.Second)
	}
}
