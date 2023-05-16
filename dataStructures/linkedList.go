package main

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	items []Node
	head  Node
	tail  Node
}

func (list *LinkedList) isEmpty() bool {
	return len(list.items) == 0
}

func (list *LinkedList) append(value int) {
	newNode := Node{data: value}
	if list.isEmpty() {
		list.head = newNode
		list.tail = list.head
	}

	newNode.next = &list.head
	list.head = newNode
	list.items = append(list.items, newNode)
}

func (list *LinkedList) prepend(value int) {
	newNode := Node{data: value}
	if list.isEmpty() {
		list.head = newNode
		list.tail = list.head
	}

	list.tail.next = &newNode
	list.tail = newNode
	list.items = append(list.items, newNode)
}
