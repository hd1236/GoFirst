package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	i    int
	lock sync.Mutex //加锁
}

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.i++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.i
}

func main() {
	var a atomicInt
	a.increment()

	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
