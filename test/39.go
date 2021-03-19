package main

import (
	"fmt"
)

func main() {
	str := "hello_世界"
	fmt.Println(str)
	// output: hello_世界

	for i := 0; i < len(str); i++ {
		fmt.Printf("[%d %c] ", str[i], str[i])
	}
	fmt.Println()
	// output: [104 h] [101 e] [108 l] [108 l] [111 o] [95 _] [228 ä] [184 ¸] [150 ] [231 ç] [149 ] [140 ]

	strBytes := []byte(str)
	fmt.Println(strBytes)
	// output: [104 101 108 108 111 95 228 184 150 231 149 140]
	for i := 0; i < len(str); i++ {
		fmt.Printf("[%d %c] ", strBytes[i], strBytes[i])
	}
	fmt.Println()
	// output: [104 h] [101 e] [108 l] [108 l] [111 o] [95 _] [228 ä] [184 ¸] [150 ] [231 ç] [149 ] [140 ]

	strRunes := []rune(str)
	fmt.Println(strRunes)
	// output: [104 101 108 108 111 95 19990 30028]
	for i := 0; i < len(strRunes); i++ {
		fmt.Printf("[%d %c] ", strRunes[i], strRunes[i])
	}
	fmt.Println()
	// output: [104 h] [101 e] [108 l] [108 l] [111 o] [95 _] [19990 世] [30028 界]

	for _, s := range str {
		fmt.Printf("[%d %c] ", s, s)
	}
	fmt.Println()
	// output: [104 h] [101 e] [108 l] [108 l] [111 o] [95 _] [19990 世] [30028 界]
}
