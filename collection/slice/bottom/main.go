package main

import (
	"fmt"
	"strconv"
)

func passSlice(slc []string) {
	slc[0] = "xiaomi"
	for i := 0; i < 10; i++ {
		slc = append(slc, strconv.Itoa(i))
	}
}

/*type slice struct {
	array unsafe.Pointer
	len int
	cap int
}*/

func main() {

	slc1 := []string{"a", "b", "c"}
	// 切片做为参数传递是值传递
	// 切片的底层是包含数组指针的结构体,这里会拷贝结构体,即这两个结构体的数组指针指向同一个数组
	passSlice(slc1)
	fmt.Println(slc1)

	slc3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slc4 := slc3[1:5]
	slc5 := slc3[2:8]
	fmt.Println(len(slc4), cap(slc4))
	fmt.Println(len(slc5), cap(slc5))
	slc5 = append(slc5, 5, 15, 25, 35, 45, 55) // 这里会扩容,slc5的数组指针会指向一个新的数组
	slc5[0] = 25
	fmt.Println(slc4)
	fmt.Println(slc5)

	var slc6 []int
	for i := 0; i < 1024; i++ {
		slc6 = append(slc6, i)
		fmt.Printf("len: %d, cap: %d\n", len(slc6), cap(slc6))
	}
}
