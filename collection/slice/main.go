package main

import "fmt"

func main() {
	var slc1 []string
	fmt.Printf("%T\n", slc1)
	if slc1 == nil {
		fmt.Println("nil")
	}

	//slc1[0] = "xbox" // 索引越界
	slc1 = append(slc1, "xbox", "nintendo")
	fmt.Println(slc1)

	// 切片的初始化
	// 1.[]string{}
	slc2 := []string{"xbox", "nintendo"}
	fmt.Println(slc2)

	// 2.截取数组
	arr := [5]string{"onepiece", "dorohedoro", "naruto", "dangdadang", "dragonball"}
	slc3 := arr[0:5] // 左闭右开
	fmt.Printf("%T\n", slc3)
	fmt.Println(slc3)

	// 3.make函数,用于切片,map,channel的创建和初始化,会分配内存
	slc4 := make([]string, 3)
	slc4[0] = "onepiece"
	slc4[1] = "naruto"
	fmt.Println(slc4)

	// 截取切片
	fmt.Println(slc3[1:]) // [1:len)
	fmt.Println(slc3[:3]) // [0:2]
	fmt.Println(slc3[:])  // [0:len)

	slc5 := []string{"xiaomi", "apple"}
	slc6 := []string{"pixel", "nokia", "samsung"}
	//slc5 = append(slc5, slc6...)
	slc5 = append(slc5, slc6[1:]...) // 拼接切片
	fmt.Println(slc5)

	// 删除某个元素
	slc7 := append(slc3[:2], slc3[3:]...)
	fmt.Println(slc7)

	// 复制切片
	slc3Copy1 := slc3[:]
	fmt.Println(slc3Copy1)

	//var slc3Copy2 []string // 不会分配内存
	slc3Copy2 := make([]string, len(slc3))
	copy(slc3Copy2, slc3)
	fmt.Println(slc3Copy2)

	slc3[3] = "berserk"
	fmt.Println(slc3Copy1)
	fmt.Println(slc3Copy2)
}
