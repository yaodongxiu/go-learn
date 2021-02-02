package main

import (
	"fmt"
	do "go-learn/interface-sample"
)

func main() {
	var aaa int
	ddd := do.BaseQuestion{}
	aaa = ddd.GetQuestionType()
	fmt.Println(aaa)
}
