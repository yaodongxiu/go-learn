package main

import (
	"fmt"
)

type FileInterface interface {
	getFileData() int
	setFileData(data int)
	write(data interface{}) error
	close()
}

type File struct {
	fileData int
}

func (f *File) getFileData() int {
	return f.fileData
}

func (f *File) setFileData(data int) {
	f.fileData = data
}

func (f *File) write(data interface{}) error {
	// 模拟写入数据
	fmt.Println("write:", data)
	return nil
}

func (f *File) close() {
	fmt.Println("closing: File stream")
}

func main() {
	// 声明一个DataWriter的接口
	var file FileInterface

	// 将接口赋值*File
	file = new(File)

	// 使用DataWriter接口进行数据写入
	err := file.write("hello world")
	if err != nil {
		fmt.Println(err.Error())
	}

	f := File{fileData: 2}
	fmt.Println(f.getFileData())
	f.setFileData(12)
	fmt.Println(f.getFileData())
	f.close()
}
