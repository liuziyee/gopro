package main

import "fmt"

func deferReturn() (i int) {
	// defer可以调用函数
	defer func() {
		i++
	}()
	return 5 // 这里可以看做i=5;return i
}

func main() {
	// defer语句会在函数返回之前调用,类似于finally
	defer fmt.Println("go")
	defer fmt.Println("java")
	fmt.Println("敲代码")

	fmt.Println(deferReturn())
}
