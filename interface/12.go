package main

import "fmt"

// define interface and design behavior here
type DataWriter interface {
	WriteData(data interface{}) error
}

type DataCloser interface {
	CloseData()
}

// define property here
type file struct {
	fileData int
}

// implement the interface you want to use
func (f *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func (f *file) CloseData() {
	f.fileData = 12
	fmt.Println(f.fileData)
	fmt.Println("closing file stream")
}
