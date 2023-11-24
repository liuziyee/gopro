package main

import (
	"fmt"
)

func main() {
	user := "root"
	password := "root"
	if user == "root" && password == "root" {
		fmt.Println("登陆成功")
	} else {
		fmt.Println("用户名或密码错误")
	}

	/*for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 3; {
		fmt.Println(i)
		i++
	}

	var j int
	for j < 10 {
		time.Sleep(time.Second * 1)
		fmt.Println(j)
		j++
	}*/

	// for-range可以取得key和value
	// 对于字符串,数组和切片,key为索引,value为对应的值的拷贝
	// 对于map,key为map的key,value为map的value的拷贝
	str := "nintendo任天堂"
	//for key, value := range str {
	//	fmt.Printf("%d, %c\n", key, value)
	//}

	for _, value := range str {
		fmt.Printf("%c\n", value)
	}

	// nintendo任天堂的字节数是17,string的底层是[]byte,索引范围是0~16,[]rune的索引范围是0~10
	runeArr := []rune(str)
	for key := range runeArr {
		fmt.Printf("%c\n", runeArr[key])
	}
}
