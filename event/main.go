package main

import "fmt"

type Actor struct {
}

func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("on event:", param)
}

// 全局事件
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func main() {
	a := new(Actor)
	RegisterEvent("OnSkill", a.OnEvent)
	RegisterEvent("OnSkill", GlobalEvent)
	CallEvent("OnSkill", 100)
}
