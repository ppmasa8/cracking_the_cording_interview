package main

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}
