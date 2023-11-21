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
	atoi, err := strconv.Atoi(str) // 字符串转整型
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(atoi)

	var in = 0
	fmt.Println(strconv.Itoa(in)) // 整型转字符串

	// parsexxx:字符串转基本类型
	f64, err := strconv.ParseFloat("3.14", 64) // 字符串转浮点数
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(f64)

	in64, err := strconv.ParseInt("30", 8, 64) // 字符串转整型
	fmt.Println(in64)

	boo, err := strconv.ParseBool("0")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(boo)

	// formatxxx:基本类型转字符串
	formatBool := strconv.FormatBool(true)
	fmt.Println(formatBool)

	fmt.Println(strconv.FormatFloat(3.1415, 'f', -1, 64))
}
