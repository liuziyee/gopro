package pkg1

import "fmt"

func Touch() Human {
	h := Human{"black"}
	fmt.Println("捏人:", h)
	return h
}

func init() {
	fmt.Println("初始化")
}
