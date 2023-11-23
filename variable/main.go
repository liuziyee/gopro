package main

import "fmt"

// 全局变量
var id = 1

func main() {
	// 先声明后赋值
	//var id int
	//id = 0

	// 声明+赋值,可以省略数据类型
	//var id = 0

	// 省略var和数据类型
	id := 0
	fmt.Println(id)

	// 有零值(默认值)
	var str string
	fmt.Println(str)

	// 声明多个变量
	var var1, var2, var3 = "xbox", 117, "nintendo"
	fmt.Println(var1, var2, var3)
}
