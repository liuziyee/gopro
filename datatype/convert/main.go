package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int8 = 3
	//var b = uint8(a)

	//var c float32 = 3.14
	//var d = int32(c)

	type ID int // 定义别名
	var e = ID(a)
	fmt.Println(e)

	var str = "12"
	atoi, err := strconv.Atoi(str) // 字符串转数字
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(atoi)

	var in = 0
	fmt.Println(strconv.Itoa(in)) // 数字转字符串
}
