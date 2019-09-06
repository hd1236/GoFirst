package main

import (
	"fmt"
	"time"
)

func printHello(i int) {
	for {
		fmt.Printf("hello from "+
			"goroutine by single func %d\n", i)
	}
}

func main() {
	//var a [10]int
	for i := 0; i < 10; i++ {
		//go func(i int) {
		//    for {
		//        fmt.Printf("hello from " +
		//           "goroutine %d",i)
		//        //a[i]++
		//        //runtime.Gosched()//交出控制权
		//    }
		//}(i)

		go printHello(i)
	}

	time.Sleep(time.Millisecond)
	//fmt.Println(a)
}
