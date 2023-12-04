package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100) // 计数器加100

	for i := 0; i < 100; i++ {
		go func(tmp int) {
			defer wg.Done() // 计数器减1
			fmt.Println(tmp)
		}(i)
	}

	wg.Wait() // 计数器为0,会唤醒阻塞的协程
	fmt.Println("子协程结束")
}
