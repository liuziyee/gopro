package main

import (
	"container/list"
	"fmt"
)

func main() {
	//var ls list.List
	ls := list.New()

	ls.PushBack("go")
	ls.PushBack("java")
	ls.PushBack("python")

	ls.PushFront("c#")

	// 插入元素
	i := ls.Front()
	for ; i != nil; i = i.Next() {
		if i.Value.(string) == "java" {
			break
		}
	}
	ls.InsertBefore("c++", i)

	for i := ls.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	for i := ls.Back(); i != nil; i = i.Prev() {
		fmt.Println(i.Value)
	}
}
