package main

import (
	"fmt"
	"strings"
)

func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case float64:
		af, _ := a.(float64)
		bf, _ := b.(float64)
		return af + bf
	case string:
		as, _ := a.(string)
		bs, _ := b.(string)
		return as + bs
	default:
		return nil
	}
}

func main() {
	a := "go "
	b := "pro"
	inf := add(a, b)
	// 类型断言
	str, _ := inf.(string)
	fmt.Println(strings.Split(str, " "))
}
