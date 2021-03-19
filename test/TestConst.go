package main

import "fmt"

type BookingType int

const (
	Default = iota
	Finish
	Cancel
)

func main() {
	fmt.Println(Default, Finish, Cancel)
	var bookingType1 BookingType = Default
	var bookingType2 BookingType = Finish
	var bookingType3 BookingType = Cancel
	fmt.Println(bookingType1)
	fmt.Println(bookingType2)
	fmt.Println(bookingType3)
}

func (bt BookingType) String() string {
	switch bt {
	case Finish:
		return "默认"
	case Cancel:
		return "落位完成"
	default:
		return "取消落位"
	}
}
