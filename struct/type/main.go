package main

import (
	"fmt"
	"strconv"
)

type MyInt2 int // 自定义类型

// 给自定义类型绑定方法
func (i MyInt2) String() string {
	return strconv.Itoa(int(i))
}

func main() {
	// type: 1.定义结构体 2.定义接口 3.定义类型别名 4.自定义类型

	// 类型别名会在编译时替换为对应的类型
	type MyInt1 = int
	var a MyInt1 = 5
	fmt.Printf("%T\n", a)

	var b MyInt2 = 5
	fmt.Printf("%T\n", b)
	fmt.Println(b.String())

	var c interface{} = "xbox"
	switch c.(type) {
	case string:
		fmt.Println("字符串")
	}
}
