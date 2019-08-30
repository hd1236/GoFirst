package main

import "fmt"

/**
函数是别称6-1
*/
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
func main() {
	//6-1
	a := adder()
	fmt.Println(a(10))
	fmt.Println(a(5))
	//for i := 0;i<10 ;i++  {
	//	fmt.Printf("0...%d的和是%d\n",i,a(i))
	//}
}
