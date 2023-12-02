// 包名最好和目录相同
package pkg1

type Human struct {
	// 首字母小写,只在包内部可见,不能被其他包访问
	// 首字母大写,可以被其他包访问
	Skin string
}
