package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name      string `json:"name"`
	AgeNumber int
	Teacher   string `json:"teacher"`
}

type Item struct {
	StudentName int `json:"name"`
	AgeNumber   int
}

func main() {
	a := []int{1, 2, 3}
	b := a
	b = nil
	fmt.Println(a, b)

	stu := Student{
		Name:      "xiao ming",
		AgeNumber: 18,
		Teacher:   "lao shi",
	}
	strJson, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(strJson))
	item := Item{}
	err = json.Unmarshal(strJson, &item)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(item)
}
