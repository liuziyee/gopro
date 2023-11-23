package main

import "fmt"

func main() {
	user := "root"
	password := "root"
	if user == "root" && password == "root" {
		fmt.Println("登陆成功")
	} else {
		fmt.Println("用户名或密码错误")
	}
}
