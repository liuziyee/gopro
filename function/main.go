package main

import "fmt"

func multiAdd(items ...int) (sum int, err error) { // 这里已经声明了sum,err这两个变量
	for _, value := range items {
		sum += value
	}
	return sum, err
}

func f1(a int, f func(int, int)) {
	f(a, 5)
}

func f2(op string) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("plus")
		}
	case "-":
		return func() {
			fmt.Println("minus")
		}
	default:
		return func() {
			fmt.Println("error")
		}
	}
}

func add(a, b int) {
	fmt.Println("into add func")
	fmt.Printf("%d\n", a+b)
}

func autoIncrement() func() int {
	i := 0
	return func() int {
		// 闭包:一个函数可以访问到外部变量
		i += 1
		return i
	}
}

func main() {
	// 函数本身可以做为变量
	addVar := multiAdd
	sum, _ := addVar(1, 3, 5)
	fmt.Println(sum)

	// 函数做为另一个函数的参数传递
	f1(1, add)

	// 函数做为另一个函数的返回值
	f2("+")()

	// 匿名函数
	anonFunc := func(a, b int) {
		fmt.Println("into anonymous func")
		fmt.Printf("%d\n", a+b)
	}
	f1(1, anonFunc)

	f1(1, func(a, b int) {
		fmt.Println("into anonymous func")
		fmt.Printf("%d\n", a+b)
	})

	f3 := autoIncrement()
	for i := 0; i < 5; i++ {
		fmt.Println(f3())
	}
	f4 := autoIncrement() // 重置i
	for i := 0; i < 5; i++ {
		fmt.Println(f4())
	}
}
