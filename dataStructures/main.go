package main

import "fmt"

func main() {
	/*q := Queue{size: 2}
	s := Stack{size: 4}

	q.enqueue(1)
	q.enqueue(13)
	fmt.Println(q.values)
	fmt.Println(q.peek())
	q.dequeue()
	fmt.Println(q.values)
	q.dequeue()
	fmt.Println(q.values)
	q.dequeue()

	s.peek()
	s.push(5)
	s.push(13)
	fmt.Println(s.values)
	fmt.Println("peek: ", s.peek())
	s.pop()
	fmt.Println(s.values)
	s.pop()
	fmt.Println(s.values)
	s.pop()*/

	l := LinkedList{}
	l.append(1)
	fmt.Println(l.items)
	l.append(2)
	fmt.Println(l.items)
	l.prepend(3)
	fmt.Println(l.items)

	for idx := range l.items {
		fmt.Print(l.items[idx].data, "->")
	}
}
