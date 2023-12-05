package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	// 无缓冲通道
	//ch = make(chan int, 0)
	// 无缓冲通道:读写操作都会阻塞
	// 有缓冲通道:
	// 通道已满,写操作会阻塞,直到有协程读取
	// 通道为空,读操作会阻塞,直到有协程写入
	//fmt.Println(<-ch)

	ch = make(chan int, 3)
	go func(ch chan int) {
		// 读取到0,3后,会阻塞
		for data := range ch {
			fmt.Println(data)
		}
		fmt.Println("读取结束")
	}(ch)

	ch <- 0
	ch <- 3
	close(ch)         // 关闭通道
	fmt.Println(<-ch) // 已关闭的通道只能读取,不能写入
	time.Sleep(time.Second * 5)
}
