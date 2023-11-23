package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "dorohedoro"
	charArr := []rune(str1)
	fmt.Println(charArr)

	//str2 := "你好 \"树先生\""
	// 反引号可以用来表示字符串
	str2 := `你好 "树先生"`
	fmt.Println(str2)

	city := "西安"
	food := "凉皮"
	drink := "冰峰"
	money := 15
	fmt.Println("城市: " + city + ", 主食: " + food + ",饮料: " + drink + ", 花费: " + strconv.Itoa(money))
	fmt.Printf("城市: %s, 主食: %s, 饮料: %s, 花费: %d\n", city, food, drink, money)
	fmt.Println(fmt.Sprintf("城市: %s, 主食: %s, 饮料: %s, 花费: %d", city, food, drink, money))

	arr := []int{0, 3, 5}
	// %v: 默认打印格式
	// %#v: 打印类型和值
	// %T: 打印类型
	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("arr: %#v\n", arr)
	fmt.Printf("arr: %T\n", arr)

	// 字符串拼接
	var strBuilder strings.Builder
	strBuilder.WriteString("城市: ")
	strBuilder.WriteString(city)
	strBuilder.WriteString(", 主食: ")
	strBuilder.WriteString(food)
	strBuilder.WriteString(", 饮料: ")
	strBuilder.WriteString(drink)
	strBuilder.WriteString(", 花费: ")
	strBuilder.WriteString(strconv.Itoa(money))

	fmt.Println(strBuilder.String())
}
