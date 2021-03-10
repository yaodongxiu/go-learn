package main

import (
	"bytes"
	"fmt"
)

func printTypeValue(slist ...interface{}) string {
	var buf bytes.Buffer
	fmt.Println(buf.Cap())
	for _, v := range slist {
		itemValueString := fmt.Sprintf("%v", v)
		var itemType string
		switch v.(type) {
		case int:
			itemType = "int"
		case string:
			itemType = "string"
		case bool:
			itemType = "bool"
		}
		buf.WriteString("value: ")
		buf.WriteString(itemValueString)
		buf.WriteString(" type: ")
		buf.WriteString(itemType)
		buf.WriteString("\n")
	}
	fmt.Println(buf.Cap())
	return buf.String()
}

func main() {
	// 将不同类型的变量通过printTypeValue()打印出来
	fmt.Println(printTypeValue(100, "str", true))
}
