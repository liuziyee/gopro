package main

import "fmt"

// 结构体类似于class
type Person struct {
	name    string
	age     int
	address string
	money   float32
}

func main() {
	// 结构体的初始化
	p1 := Person{"lee", 30, "hongkong", 10000.0}
	p2 := Person{
		name:  "lee",
		money: 10000.0,
	}
	var p3 Person
	p3.money = 10.0
	fmt.Println(p3.age)

	var ps1 []Person
	ps1 = append(ps1, p1, p2)
	ps1 = append(ps1, Person{
		name: "lee",
	})

	ps2 := []Person{
		{"lee", 30, "hongkong", 10000.0},
		{
			name:  "lee",
			money: 10000.0,
		},
	}
	fmt.Println(ps2)

	// 匿名结构体
	address := struct {
		province string
		city     string
		area     string
	}{"陕西", "西安", "雁塔区"} // 实例化

	fmt.Println(address.city)
}
