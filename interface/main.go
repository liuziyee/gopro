package main

import "fmt"

type Duck interface {
	Quack()
	Shit() string
}

type Bird interface {
	Fly()
}

type Eater interface {
	Eat()
}

type MallardDuck struct {
	wing  int
	Eater // 匿名接口,即实现了Eater接口
}

type LettuceEater struct{}

type AppleEater struct{}

func (m *MallardDuck) Quack() {
	fmt.Println("gaga")
}

func (m *MallardDuck) Shit() string {
	return "shit"
}

func (m *MallardDuck) Fly() {
	fmt.Println("fly")
}

func (le *LettuceEater) Eat() {
	fmt.Println("eat lettuce")
}

func (ae *AppleEater) Eat() {
	fmt.Println("eat apple")
}

func log(items ...interface{}) {
	for _, value := range items {
		fmt.Println(value)
	}
}

func main() {
	// 空接口可以接收任意类型
	//var md any = &MallardDuck{0, &LettuceEater{}}

	// 这里类似于向上转型,即父类引用指向子类对象
	var d1 Duck = &MallardDuck{0, &LettuceEater{}}
	d1.Quack()

	var d2 Bird = &MallardDuck{0, &AppleEater{}}
	d2.Fly()

	var d3 Eater = &MallardDuck{0, &AppleEater{}}
	d3.Eat()

	slc1 := []interface{}{0, "zero"}
	log(slc1...)

	// 字符串切片和空接口切片不是同一个类型
	slc2 := []string{"0", "zero"}
	//log(slc2...)

	var slc3 []interface{}
	for _, value := range slc2 {
		slc3 = append(slc3, value)
	}
	log(slc3...)
}
