package main

import (
	"fmt"
	"sync"
)

var n int32
var wg sync.WaitGroup
var lock sync.Mutex

func plus() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		n += 1
		lock.Unlock()
		//atomic.AddInt32(&n, 1)
	}
}

func minus() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock()
		n -= 1
		lock.Unlock()
		//atomic.AddInt32(&n, -1)
	}
}

func main() {
	wg.Add(2)
	go plus()
	go minus()
	wg.Wait()
	fmt.Println(n)
}
