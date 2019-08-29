package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//包作用域
var aa = 5
var bb = 9
var tt = true

func variableTypeDedution() {
	var a, b, c, s = 3, 4, true, "dad"
	fmt.Println(a, b, c, s)
}

func variableSimple() {
	a, b, c, d := 3, 4, true, "adf"
	fmt.Println(a, b, c, d)
}

func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))

	var c = cmplx.Exp(1i*math.Pi) + 1
	fmt.Printf("%.3f\n", c)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) //强制类型转换
	fmt.Println(c)
}
func consts() {
	//const filename = "abc.txt"
	//const a,b = 3,4
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)

}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	//b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func grade(score int) string {
	var g string
	switch {
	case score < 60:
		g = "f"
	case score < 80:
		g = "c"
	case score < 90:
		g = "b"
	case score <= 100:
		g = "a"
	default:
		panic(fmt.Sprint("hehe"))
	}
	fmt.Println(g)
	return g
}

//函数式编程,函数的参数也可以是函数
func apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}

//可变参数
func sum(numbers ...int) int {
	var sum int
	for j, i := range numbers {
		fmt.Printf("j=%d,i=%d\n", j, i)
		sum += i
	}
	return sum
}

//指针，go语言只有值传递
func swap(a, b *int) {
	a, b = b, a
}
func swap1(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	//fmt.Println("Helo World")
	//
	//variableTypeDedution()
	//variableSimple()
	//fmt.Println(aa,bb,tt)
	//
	//euler()
	//
	////2-2
	//triangle()
	//2-3
	//consts()
	//enums()

	//2-4
	//grade(100)
	//grade(40)
	//grade(70)
	//grade(101)

	//2-5\函数\可变参数列表
	//ret := apply(4,5, func(a int, b int) int {
	//	return a * b
	//})
	//fmt.Println("result:",ret)
	//
	//fmt.Println(sum(1,9,3,6,5))

	a, b := 3, 4
	//swap(&a,&b)
	swap1(&a, &b)
	fmt.Println("a,b:", a, b)
	var c int = 5
	d := &c
	var f *int
	f = &*d
	fmt.Println("指针：", *d, *f)

}
