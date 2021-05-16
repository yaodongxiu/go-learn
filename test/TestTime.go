package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time
	tu := t.Unix()
	fmt.Println(tu)
	//var tp *time.Time
	//tpu := tp.Unix()
	//fmt.Println(tpu) // error
}
