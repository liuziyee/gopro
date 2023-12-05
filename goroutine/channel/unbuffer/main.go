package main

import (
	"fmt"
	"time"
)

var num, ltr = make(chan bool), make(chan bool) // 无缓冲通道,可用于协程间的通信

func printNum() {
	i := 0
	for {
		<-num // 读通道会阻塞,直到有协程写入
		fmt.Printf("%d%d", i, i+1)
		i += 2
		ltr <- true // 写通道会阻塞,直到有协程读取
	}
}

func printLtr() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		<-ltr
		fmt.Printf("%c%c", str[i], str[i+1])
		if i == len(str)-2 {
			return
		}
		i += 2
		num <- true
	}
}

func main() {
	go printNum()
	go printLtr()
	num <- true
	time.Sleep(time.Second * 5)
}
