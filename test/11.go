package main

import "fmt"

func main() {
	map1 := make(map[int8]float32)
	map1[0] = 1.0
	map1[1] = 1.1
	for key, value := range map1 {
		fmt.Println(fmt.Sprintf("key[%d] => value[%.2f]", key, value))
	}
}
