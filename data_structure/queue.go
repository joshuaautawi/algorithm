package datastructure

import "fmt"

type Node2 struct {
	Value int
	Next  *Node2
}
type Queue struct {
	Head *Node2
	Tail *Node2
	Size int
}

func (q *Queue) AddOne(node *Node2) {
	q.Size++
	if q.Tail == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
}

func (q *Queue) RemoveOne() {
	if q.Head == nil {
		return
	}
	q.Size--
	head := q.Head
	q.Head = q.Head.Next
	head.Next = nil

}

func Run() {
	node1 := Node2{Value: 1}
	node2 := Node2{Value: 2}
	node3 := Node2{Value: 3}
	data := Queue{}
	data.AddOne(&node1)
	data.AddOne(&node2)
	data.AddOne(&node3)
	data.RemoveOne()
	data.Print()
}

func (q *Queue) Print() {
	fmt.Println("Start Function ")
	node := q.Head
	for node.Next != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
	fmt.Println(node.Value)
}
