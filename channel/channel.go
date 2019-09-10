package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	//for{
	//    n,ok :=<-c
	//    if !ok{
	//       break
	//    }
	//    fmt.Printf("Worker %d received %c\n",id,n)
	//}
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

//channel缓冲
func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

//关闭channel
func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	//普通channel，一等公民，作为参数，返回值
	chanDemo()

	//buffered
	//bufferedChannel()

	//channel发送方close，接收方判断
	//channelClose()
}
