package main

import "fmt"

var ans []string

func dikaerji(target [][]string) []string {
	l := len(target)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return target[0]
	}
	ans = make([]string, 0)
	backtrack(target, 0, "")
	return ans
}

func backtrack(target [][]string, index int, item string) {
	if len(item) == len(target) {
		ans = append(ans, item)
	} else {
		for _, v := range target[index] {
			item += v
			backtrack(target, index+1, item)
			item = item[:len(item)-1]
		}
	}
}

func main() {
	var target [][]string
	target = [][]string{{"A", "B", "C"}, {"+", "-", "*", "/"}, {"1", "2", "3"}}
	fmt.Println(dikaerji(target))
}
