package main

import (
	Queue "AoLiGei/Object"

	"fmt"
)

func main() {
	q := Queue.Queue{1}

	q.Push(3)
	q.Push(5)
	pop := q.Pop()
	pop1 := q.Pop()
	pop2 := q.Pop()
	fmt.Println(q, pop, pop1, pop2)
	fmt.Println(q.IsEmpty())

	q.Push(1)
	fmt.Println(q)
}
