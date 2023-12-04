package main

import (
	"fmt"
	"time"
)

func producer(in chan<- int) {
	for i := 0; i < 10; i++ {
		in <- i * i
	}
	close(in)
}

func consumer(out <-chan int) {
	for data := range out {
		fmt.Println(data)
	}
}

func main() {
	ch := make(chan int, 3)
	// 只能写入的通道
	//var wch chan<- int = ch
	// 只能读取的通道
	//var rch <-chan int = ch

	go producer(ch)
	go consumer(ch)
	time.Sleep(time.Second * 5)
}
