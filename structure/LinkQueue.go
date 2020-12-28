package main

import "fmt"

type QueueNode struct {
	data int
	next *QueueNode
}

type LinkQueue struct {
	head   *QueueNode
	rear   *QueueNode
	length int
}

type QueueInterface interface {
	Push(data int)
	Pop() int
	Clear()
}

func (queue *LinkQueue) Push(data int) {
	queue.length++
	node := new(QueueNode)
	node.data = data
	node.next = nil
	if queue.head == nil {
		queue.head = node
		queue.rear = node
	}
	queue.rear.next = node
	queue.rear = node

	fmt.Println(fmt.Sprintf("queue push data => %d", data))
}

func (queue *LinkQueue) Pop() int {
	if queue.length <= 0 {
		return -999
	}
	queue.length--
	popData := queue.head.data
	queue.head = queue.head.next
	return popData
}

func (queue *LinkQueue) Clear() {
	for {
		popData := queue.Pop()
		if popData == -999 {
			break
		} else {
			fmt.Println(fmt.Sprintf("queue pop data => %d", popData))
		}
	}
}

func createNewLinkQueue() *LinkQueue {
	return &LinkQueue{
		nil,
		nil,
		0,
	}
}

func main() {
	var queue QueueInterface
	queue = createNewLinkQueue()
	for i := 1; i <= 10; i++ {
		queue.Push(i)
	}
	queue.Clear()
}
