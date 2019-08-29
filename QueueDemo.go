package main

import (
	"HelloWorld/Queue"
	"fmt"
)

func main() {

	q := Queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q)
	q.Pop()
	q.Pop()
	fmt.Println(q)
	fmt.Println(q.IsEmpty())
	q.Pop()
	fmt.Println(q.IsEmpty())
	v := q.Pop()
	fmt.Println(v)
	//4-4
}
