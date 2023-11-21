package main

import "fmt"

func main() {
	//var a int8
	//var b int16
	//var c int32
	//var d int64
	//var ua uint8
	//var ub uint16
	//var uc uint32
	//var ud uint64
	//var e int // 动态类型
	//
	//a = int8(b)
	//
	//var e float32
	//var f float64

	// byte和rune用来存储字符,rune可以存储中文字符
	var char1 byte
	char1 = 'x'
	char2 := char1 + 1
	fmt.Println(char1, char2)
	fmt.Printf("%c, %c\n", char1, char2)

	var char3 rune
	char3 = '宿'
	fmt.Printf("%c", char3)
}
