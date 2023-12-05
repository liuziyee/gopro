package main

import (
	"fmt"
	"time"
)

func g1(ch chan struct{}) {
	time.Sleep(time.Second)
	ch <- struct{}{}
}

func g2(ch chan struct{}) {
	time.Sleep(time.Second * 3)
	ch <- struct{}{}
}

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go g1(ch1)
	go g2(ch2)

	timer := time.NewTimer(time.Second * 5) // 5秒之后timer.C通道就绪
	for {
		// select可用于监听通道的操作是否就绪,如果没有通道(分支)就绪,会阻塞
		// 某个分支就绪,执行该分支
		// 多个分支同时就绪,随机执行某个分支
		// 无分支就绪时,会执行default代码块,select就不会阻塞
		select {
		case <-ch1:
			fmt.Println("协程1结束")
		case <-ch2:
			fmt.Println("协程2结束")
		case <-timer.C:
			fmt.Println("已超时")
			return
		default:
			time.Sleep(time.Millisecond * 500)
			fmt.Println("无通道就绪")
		}
	}
}
