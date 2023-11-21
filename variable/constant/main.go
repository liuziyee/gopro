package main

import "fmt"

func main() {
	// 常量的数据类型:布尔,数值(整数,浮点数,复数),字符串
	const PI float32 = 3.14

	// 常量组
	const (
		a = 0
		b
		x = "xbox"
		y
	)
	fmt.Println(a, b, x, y)
}
