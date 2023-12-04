package main

import (
	"fmt"
	"time"
)

// 主协程
func main() {
	/*for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i) // 闭包
		}()
	}*/

	/*for i := 0; i < 10; i++ {
		tmp := i
		go func() {
			fmt.Println(tmp)
		}()
	}*/

	for i := 0; i < 10; i++ {
		// 创建协程(协程的调度时间是无法控制的)
		go func(tmp int) {
			fmt.Println(tmp) // i的副本
		}(i)
	}

	time.Sleep(time.Second * 10)
}
