package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
var refCount int = 0

func fooFunc(i int, c chan int) {
	lock.Lock()
	refCount += 1
	c <- refCount
	lock.Unlock()
	fmt.Printf("fooFunc %d\n", i)
}

func main() {
	c := make(chan int, 100)

	nrCoroutine := 10
	for i := 0; i < nrCoroutine; i++ {
		go fooFunc(i, c)
	}

	time.Sleep(time.Second)
	for i := range c {
		fmt.Printf("%d ", i)
		if refCount == nrCoroutine {
			close(c)
			break
		}
	}
	fmt.Printf("refCount = %d\n", refCount)
}
