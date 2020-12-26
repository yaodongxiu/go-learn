package main

import "fmt"

type StackNode struct {
	data int
	next *StackNode
}

type LinkStack struct {
	stackTop StackNode
	length   int
}

type StackInterface interface {
	Pop() int
	Push(data int)
	Clear()
}

func (stack *LinkStack) Push(data int) {
	stack.length++

	node := StackNode{}
	node.data = data
	node.next = &stack.stackTop
	stack.stackTop = node

	fmt.Println(fmt.Sprintf("push data => %d", data))
}

func (stack *LinkStack) Pop() int {
	if stack.length <= 0 {
		return -999
	}

	stack.length--
	popData := stack.stackTop.data
	stack.stackTop = *stack.stackTop.next
	return popData
}

func (stack *LinkStack) Clear() {
	for {
		popData := stack.Pop()
		if popData == -999 {
			break
		} else {
			fmt.Println(fmt.Sprintf("pop data => %d", popData))
		}
	}
}

func createLinkStack() *LinkStack {
	return &LinkStack{
		StackNode{
			0,
			nil,
		},
		0,
	}
}

func main() {
	var stack StackInterface
	stack = createLinkStack()

	stack.Push(1)
	stack.Clear()
}
