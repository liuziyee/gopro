package main

import "fmt"

type Mouth interface {
	Eat()
}

type Eye interface {
	Look()
}

type Bird interface {
	// Bird继承了Mouth和Eye,也即继承了Eat和Look方法
	Mouth
	Eye
	Fly()
}

type Duck struct{}

func (d Duck) Eat() {
	fmt.Println("eat")
}

func (d Duck) Look() {
	fmt.Println("look")
}

func (d *Duck) Fly() {
	fmt.Println("fly")
}

func main() {
	// 不用父接口接收的话,是否对结构体取址以及调用方法是否为指针方法都没有限制
	var d1 = &Duck{}
	d1.Fly()

	// 用父接口接收的话,已实现的接口方法都是值方法的话,是否取址没有限制
	// 已实现的接口方法有至少一个指针方法的话,要取址
	var d2 Bird = &Duck{}
	d2.Eat()
}
