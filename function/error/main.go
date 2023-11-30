package main

import (
	"errors"
	"fmt"
)

func f1() (int, error) {
	// panic会使程序退出
	//panic("this is a panic")

	// recover要放在defer语句使用,同时recover生效(即捕获到panic)的前提是defer是在panic之前定义的
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到panic: ", r)
		}
	}()

	var map1 map[string]string
	map1["0"] = "zero"

	// panic之后的代码不会执行
	defer fmt.Println("defer after panic")

	return -1, errors.New("this is an error")
}

func main() {
	if _, err := f1(); err != nil {
		fmt.Println(err)
	}
}
