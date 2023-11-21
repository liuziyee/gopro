package main

import "fmt"

func get() (int, bool) {
	return 0, true
}

func main() {
	// 匿名变量
	var _ int

	// 匿名变量不会强制使用,可以做为占位符
	_, code := get()
	if code {
		fmt.Println(code)
	}

	// 代码块
	{
		blockVar := "xbox"
		fmt.Println(blockVar)
	}
}
