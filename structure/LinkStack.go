package main

import "fmt"

type StackNode struct {
	data int
	next *StackNode
}

type LinkStack struct {
	stackTop *StackNode
	length   int
}

type StackInterface interface {
	Push(data int)
	Pop() int
	Clear()
}

func (stack *LinkStack) Push(data int) {
	stack.length++

	node := StackNode{}
	node.data = data
	node.next = stack.stackTop
	stack.stackTop = &node

	fmt.Println(fmt.Sprintf("stack push data => %d", data))
}

func (stack *LinkStack) Pop() int {
	if stack.length <= 0 {
		return -999
	}

	stack.length--
	popData := stack.stackTop.data
	stack.stackTop = stack.stackTop.next
	return popData
}

func (stack *LinkStack) Clear() {
	for {
		popData := stack.Pop()
		if popData == -999 {
			break
		} else {
			fmt.Println(fmt.Sprintf("stack pop data => %d", popData))
		}
	}
}

func createLinkStack() *LinkStack {
	return &LinkStack{
		nil,
		0,
	}
}

func main() {
	var stack StackInterface
	stack = createLinkStack()
	for i := 1; i <= 10; i++ {
		stack.Push(i)
	}
	stack.Clear()
}
