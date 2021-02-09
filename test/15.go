package main

import (
	"bytes"
	"fmt"
)

func printTypeValue(slist ...interface{}) string {
	var buf bytes.Buffer
	fmt.Println(buf.Cap())
	for _, v := range slist {
		str := fmt.Sprintf("%v", v)
		var tye string
		switch v.(type) {
		case int:
			tye = "int"
		case string:
			tye = "string"
		case bool:
			tye = "bool"
		}
		buf.WriteString("value: ")
		buf.WriteString(str)
		buf.WriteString(" type: ")
		buf.WriteString(tye)
		buf.WriteString("\n")
	}
	fmt.Println(buf.Cap())
	return buf.String()
}

func main() {
	// 将不同类型的变量通过printTypeValue()打印出来
	fmt.Println(printTypeValue(100, "str", true))
}
