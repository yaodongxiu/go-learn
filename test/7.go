package main

import "fmt"

func main() {
	theme := "狙击 start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("%c %d\n", theme[i], theme[i])
	}
	for index, s := range theme {
		fmt.Printf("%d %c %d\n", index, s, s)
	}
}
