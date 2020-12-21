package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main() {
	head := new(Node)
	head.data = 0
	//addToHead(head)
	addToTail(head)
	showNode(head)
}

func showNode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func addToHead(head *Node) *Node {
	for i := 1; i < 10; i++ {
		node := Node{data: i}
		node.next = head.next
		head.next = &node
	}
	return head
}

func addToTail(head *Node) *Node {
	tail := head
	for i := 1; i < 10; i++ {
		node := Node{data: i}
		tail.next = &node
		tail = &node
	}
	return head
}
