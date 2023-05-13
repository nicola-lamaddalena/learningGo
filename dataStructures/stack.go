package main

import "fmt"

type Stack struct {
	values []int
	size   int
}

func (s *Stack) push(element int) {
	if len(s.values) == s.size {
		fmt.Println("The stack is full")
		return
	}
	s.values = append(s.values, element)
}

func (s *Stack) pop() int {
	if len(s.values) == 0 {
		fmt.Println("The stack is empty")
		return 0
	}
	index := len(s.values) - 1
	s.values = s.values[:len(s.values)-1]
	return index
}

func (s *Stack) peek() int {
	if len(s.values) == 0 {
		fmt.Println("The stack is empty")
		return 0
	}
	index := len(s.values) - 1
	return s.values[index]
}
