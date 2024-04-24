package datastructure

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func LL() {
	node1 := Node{Data: 1, Next: nil}
	node2 := Node{Data: 2, Next: nil}
	node3 := Node{Data: 3, Next: nil}
	node4 := Node{Data: 4, Next: nil}
	link := LinkedList{Head: &node1}
	link.append(&node2)
	link.append(&node3)
	link.prepend(&node4)
	link.Print()

}

func (l *LinkedList) append(newNode *Node) {
	node := l.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = newNode
}

func (l *LinkedList) prepend(newNode *Node) {
	node := l.Head
	l.Head = newNode
	l.Head.Next = node
}

func (l *LinkedList) Print() {
	node := l.Head
	for node.Next != nil {
		fmt.Println(node.Data)
		node = node.Next
	}
	fmt.Println(node.Data)
}
