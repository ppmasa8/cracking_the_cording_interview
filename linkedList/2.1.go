package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	val  int
	next *Node
}

type ListNode struct {
	head *Node
}

func newNode(x int, cp *Node) *Node {
	newcp := new(Node)
	newcp.val, newcp.next = x, cp
	return newcp
}

func newListNode() *ListNode {
	lst := new(ListNode)
	lst.head = new(Node)
	return lst
}

func (lst *ListNode) Insert(n int) {
	cp := lst.head
	cp.next = newNode(n, cp.next)
}

func main() {
	node := newListNode()
	for i := 0; i < 5; i++ {
		n := rand.Intn(100)
		fmt.Println(n)
		node.Insert(n)
	}


}
