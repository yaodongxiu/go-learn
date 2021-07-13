package main

import "fmt"

type BookingType int

const (
	Default BookingType = iota
	Finish
	Cancel
)

var BookingTypeMap = map[BookingType]string{
	Default: "默认",
	Finish:  "落位完成",
	Cancel:  "取消落位",
}

func main() {
	bookingType1 := Finish
	fmt.Println(bookingType1.String())
}

func (bt BookingType) String() string {
	if msg, ok := BookingTypeMap[bt]; ok {
		return msg
	}
	return ""
}
