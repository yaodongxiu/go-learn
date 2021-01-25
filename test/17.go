package main

import (
	"encoding/json"
	"fmt"
)

type Screen struct {
	Size       float32 //屏幕尺寸
	ResX, ResY int     //水平和垂直分辨率

}

type Battery struct {
	Capacity int //电池容量

}

func getJsonData() []byte {

	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery{
			2910,
		},
		HasTouchID: true,
	}

	jsonData, _ := json.Marshal(raw)

	return jsonData

}
func main() {

	jsonData := getJsonData()

	fmt.Printf(string(jsonData))

	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	err := json.Unmarshal(jsonData, &screenAndTouch)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%+v\n", screenAndTouch)
}
