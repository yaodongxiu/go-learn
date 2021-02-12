package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	tip1 := "genji is a ninja"
	fmt.Println(len(tip1))
	fmt.Println(utf8.RuneCountInString(tip1))
	tip2 := "忍者"
	fmt.Println(len(tip2))
	fmt.Println(utf8.RuneCountInString("忍者"))
	fmt.Println(utf8.RuneCountInString("龙忍出鞘,fight!"))
}
