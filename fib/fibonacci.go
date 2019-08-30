package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/**
函数是编程6-2
*/
func fibonacci() intGet {
	sum1 := 0
	sum2 := 1
	return func() int {
		//var temp = sum2
		//sum2 = sum1 +temp
		//sum1 = temp
		sum1, sum2 = sum2, sum1+sum2 //可以双赋值
		//fmt.Println(sum1)
		return sum1
	}
}

type intGet func() int //定义函数类型，实现接口

func (g intGet) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	f := fibonacci()
	//f()
	//f()
	//f()
	//f()
	//f()
	//f()
	//f()
	printFileContents(f)
}
