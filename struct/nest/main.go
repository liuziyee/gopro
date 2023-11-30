package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
	money   float32
}

type Student1 struct {
	// 嵌套方式1
	p      Person
	school string
}

type Student2 struct {
	// 嵌套方式2:匿名嵌套
	Person
	school  string
	address string
}

func main() {
	s1 := Student1{
		Person{"lee", 30, "hongkong", 10000.0},
		"hongkong",
	}
	fmt.Println(s1.p.address)

	s2 := Student2{
		Person{"lee", 30, "hongkong", 10000.0},
		"hongkong",
		"hongkong",
	}
	s2.address = "台湾" // 这里访问的是外部的address
	fmt.Println(s2)
}
