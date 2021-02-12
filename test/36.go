package main

import "fmt"

func main() {
	defer fmt.Println("in main111")
	defer fmt.Println("in main222")
	defer func() {
		fmt.Println("in main333")
		defer fmt.Println("in main444")
		defer func() {
			fmt.Println("in main555")
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")

	// Output:
	//	in main333
	//	in main555
	//	in main444
	//	in main222
	//	in main111
	//panic: panic once
	//panic: panic again
	//panic: panic again and again
	//
	//	goroutine 1 [running]:...

}
