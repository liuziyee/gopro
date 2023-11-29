package main

import "fmt"

func main() {
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("map1 is nil")
	}
	//map1["0"] = "zero" // map为nil,put会报错
	map2 := map[string]string{}
	if map2 == nil {
		fmt.Println("map2 is nil")
	}
	map2["0"] = "zero"

	// map的初始化
	//1.map[key]value{}
	map3 := map[string]string{
		"0": "zero",
		"1": "one",
		"2": "two",
	}

	fmt.Println(map3["0"])
	map3["3"] = "three"
	fmt.Println(map3)

	// 2.make函数
	map4 := make(map[string]string, 3)
	map4["0"] = "zero"
	fmt.Println(map4)

	// 遍历map
	for key, value := range map3 {
		fmt.Println(key, value)
	}

	for key := range map3 {
		fmt.Println(key, map3[key])
	}

	// if的另一种写法
	if _, ok := map3["-1"]; ok {
		fmt.Printf("in")
	} else {
		fmt.Println("not in")
	}

	// 删除元素
	delete(map3, "0")
	fmt.Println(map3)
}
