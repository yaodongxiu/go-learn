package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("a")
	l.PushBack(222)
	l.PushBack(3.33)

	for element := l.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

}
