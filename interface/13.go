package main

import (
	"fmt"
)

func main() {
	// 声明一个DataWriter的接口
	var writer DataWriter
	// 将接口赋值*file
	writer = new(file)
	// 使用DataWriter接口进行数据写入
	err := writer.WriteData("hello world")
	if err != nil {
		fmt.Println(err.Error())
	}

	//var f DataCloser
	//f = new(file)
	f := file{fileData: 2}
	f.CloseData()
	fmt.Println(f.fileData)

}
