package main

import (
	"fmt"
	"time"
)

func main() {
	dd := make(chan int, 1)
	close(dd) // chan close 后，不再有阻塞的效果，而且无论chan是有缓冲或者没有缓冲的，close后，vv=0， ok=false
	vv, ok := <-dd
	fmt.Println(ok)   // false
	fmt.Println(vv)   // 0
	fmt.Println(<-dd) // 0
	fmt.Println(<-dd) // 0
	fmt.Println(<-dd) // 0
	cc := make(chan int, 1)
	i := 0
	for {
		select {
		case v := <-cc:
			fmt.Println(v)
		case <-time.After(2 * time.Second):
			fmt.Println("haha")
			i++
			cc <- i
		}
	}
}
