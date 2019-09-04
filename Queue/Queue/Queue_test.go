package Queue

import "fmt"

/**
做文档的example，既能做测试，也能做文档
*/
func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	//Output:
	//1
	//2
	//false
	//3
	//true
}
