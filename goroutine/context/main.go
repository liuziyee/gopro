package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 通道版
/*func cpu(close chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-close:
			fmt.Println("程序退出")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("CPU占用35%")
		}
	}
}*/

func cpu(ctx context.Context) {
	fmt.Printf("id: %d\n", ctx.Value("id"))
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("停止打印")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("CPU占用35%")
		}
	}
}

func main() {
	// 通道版
	/*wg.Add(1)
	var close = make(chan struct{})
	go cpu(close)
	time.Sleep(time.Second * 5)
	close <- struct{}{}

	wg.Wait()
	fmt.Println("程序退出")*/

	/*wg.Add(1)
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, _ := context.WithCancel(ctx1)
	go cpu(ctx2)
	time.Sleep(time.Second * 5)
	cancel1()

	wg.Wait()
	fmt.Println("程序退出")*/

	/*wg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	go cpu(ctx)

	wg.Wait()
	fmt.Println("程序退出")*/

	wg.Add(1)
	ctx1, _ := context.WithTimeout(context.Background(), time.Second*5)
	ctx2 := context.WithValue(ctx1, "id", 0)
	go cpu(ctx2)

	wg.Wait()
	fmt.Println("程序退出")
}
