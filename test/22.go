package main

import (
	. "fmt"
	do "go-learn/interface-sample"
)

func main() {
	a := do.BaseQuestion{
		QuestionId:      1,
		QuestionContent: "22",
		QuestionType:    333,
	}
	Println(a)
}
