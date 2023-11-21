package main

import "fmt"

func main() {
	// iota:用于生成常量的递增计数器
	// 每次使用iota,都会自动递增,自增类型默认为int
	const (
		ERR1 = iota
		ERR2
		ERR3 = "xbox" // iota仍然会递增,下同
		ERR4
		ERR5
		ERR6 = 3
		ERR7 = iota
	)
	fmt.Println(ERR1, ERR2, ERR3, ERR7)

	// iota在每个const块里都会归零
	const (
		ERR0 = iota
	)
	fmt.Println(ERR0)
}
