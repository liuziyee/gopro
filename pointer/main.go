package main

import "fmt"

type Person struct {
	money float32
}

func money(p *Person) {
	p.money = 10.0
}

func main() {
	p1 := Person{10000.0}
	money(&p1)
	var p2 *Person = &p1
	fmt.Printf("%p\n", p2)

	// nil指针
	var p3 *Person
	if p3 == nil {
		fmt.Println("nil指针")
	}
	//fmt.Println(p3.money)
	p3 = &Person{10000.0}
	p3.money = 10.0

	// 指针的初始化
	// 1. &取址
	p4 := &Person{}
	fmt.Println(p4.money)

	var p5 Person // 结构体可以不做初始化,即字段是零值
	p6 := &p5
	fmt.Println(p6.money)

	// 2.new函数
	p7 := new(Person)
	fmt.Println(p7.money)

	// 初始化的两个函数
	// make函数:切片,map,channel
	// new函数:指针
	// 切片,map,指针要初始化,结构体可以不做初始化
}
