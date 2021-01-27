package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64
	wg0     sync.WaitGroup
	mutex   sync.Mutex
)

func incCounter() {
	defer wg0.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		fmt.Println(666)
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
			//counter++
		}
		mutex.Unlock()
	}
}

func main() {
	wg0.Add(2)

	go incCounter()
	go incCounter()

	wg0.Wait()
	fmt.Println(counter)
}
