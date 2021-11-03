package main
/*** 方針 ***/
// ソートされていない連結リストから、重複する要素を削除するには、要素を配列に追加していき、
// 要素がかぶった場合には、削除してforを回していく。

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

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func main() {
	node := newListNode()
	for i := 0; i < 3; i++ {
		n := rand.Intn(100)
		fmt.Println(n)
		node.Insert(n)
	}
	// 重複させる用
	fmt.Println(10)
	fmt.Println(10)
	node.Insert(10)
	node.Insert(10)

    arr := []int{}

	cp := newListNode()

	for node.head != nil {
		if contains(arr, node.head.val) {
			cp.head.next = node.head.next
		} else {
			arr = append(arr, node.head.val)
			cp = node
		}
		node.head = node.head.next
	}
}
// 時間計算量O(n)