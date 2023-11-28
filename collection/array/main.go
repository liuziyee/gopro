package main

import "fmt"

func main() {
	/*var arr1 [3]string
	arr1[0] = "go"
	arr1[1] = "java"
	arr1[2] = "c#"
	fmt.Printf("%T\n", arr1)*/

	arr1 := [3]string{"go", "java", "c#"}
	for _, value := range arr1 {
		fmt.Println(value)
	}

	// 通过数组的某个(多个)索引赋值
	arr2 := [3]string{1: "go"}
	for _, value := range arr2 {
		fmt.Println(value)
	}

	// 缺省数组长度
	arr3 := [...]string{"go", "java", "c#"}
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	arr4 := [3]string{"go", "c#", "java"}
	// 只能比较同类型数组
	if arr3 == arr4 {
		fmt.Println("equal")
	}

	// 二维数组
	var arr5 [3][3]string
	arr5[0] = [3]string{"go", "1h", "goland"}
	arr5[1] = [3]string{"java", "3h", "idea"}
	arr5[2] = [3]string{"c#", "5h", "vscode"}

	for i := 0; i < len(arr5); i++ {
		for j := 0; j < len(arr5[i]); j++ {
			fmt.Print(arr5[i][j] + " ")
		}
		fmt.Println()
	}

	for _, row := range arr5 {
		for _, value := range row {
			fmt.Print(value + " ")
		}
		fmt.Println()
	}

}
