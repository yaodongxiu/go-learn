package main

import (
	"fmt"
)

func main() {
	str := "C语言中文网\nGo语言教程"
	fmt.Println(str)
	str = "123"
	fmt.Println(str)

	aaa := "asdf中文"
	var bytess []byte
	bytess = []byte(aaa)
	fmt.Println(bytess)
	fmt.Println(string(bytess))
	var runess []rune
	runess = []rune(aaa)
	fmt.Println(runess)
	fmt.Println(string(runess))

	sss1, err := fmt.Printf("value => %s", "sss_value")
	sss2 := fmt.Sprintf("value => %s", "sss_value")
	fmt.Println(sss1, err, sss2)

}
