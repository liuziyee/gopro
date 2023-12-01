package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
	money   float32
}

type Student struct {
	Person
	school string
}

// 绑定方法到结构体
// 值接收器(值传递)
func (p Person) toString1() {
	fmt.Printf("%s, %d, %s, %f\n", p.name, p.age, p.address, p.money)
}

// 指针接收器(引用传递)
func (p *Person) toString2() {
	p.money = 10.0
	fmt.Printf("%s, %d, %s, %f\n", p.name, p.age, p.address, p.money)
}

func main() {
	p1 := Person{
		name:    "lee",
		age:     30,
		address: "hongkong",
		money:   10000.0,
	}
	p1.toString2() // 值可以调用指针方法
	fmt.Println(p1)

	s1 := Student{
		Person{"lee", 30, "hongkong", 10000.0},
		"hongkong",
	}
	s1.toString1()
}
