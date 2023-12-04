package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var rwLock sync.RWMutex // 读写锁
	wg.Add(6)

	go func() {
		time.Sleep(time.Second)
		defer wg.Done()
		rwLock.Lock() // 写锁
		defer rwLock.Unlock()
		fmt.Println("取得写锁")
		time.Sleep(time.Second * 5)
	}()

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			for {
				rwLock.RLock() // 读锁
				fmt.Println("取得读锁")
				time.Sleep(time.Millisecond * 500)
				rwLock.RUnlock()
			}
		}()
	}

	wg.Wait()
}
