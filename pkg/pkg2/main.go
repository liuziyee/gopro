package main

// 普通导入,会在使用该包的地方调用其init函数
// 包别名 包路径
import human "gopro/pkg/pkg1"

// 匿名导入,程序启动时会调用该包的init函数
//import _ "gopro/pkg/pkg1"

func main() {
	h := human.Human{Skin: "black"}
	h.Skin = "yellow"
	human.Touch()
}
