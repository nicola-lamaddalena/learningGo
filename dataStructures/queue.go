package main

import "fmt"

type Queue struct {
	values []int
	size   int
}

func (q *Queue) enqueue(element int) {
	if len(q.values) == q.size {
		fmt.Println("The queue is full")
		return
	}
	q.values = append(q.values, element)
}

func (q *Queue) dequeue() {
	if len(q.values) == 0 {
		fmt.Println("The queue is empty")
		return
	}
	q.values = q.values[1:]
}

func (q *Queue) peek() int {
	if len(q.values) == 0 {
		fmt.Println("The queue is empty")
		return 0
	}
	return q.values[len(q.values)-1]
}
