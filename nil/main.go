package main

import "fmt"

type Person struct {
	money float32
}

func main() {
	/**
	类型的零值
	bool false
	int,float 0
	string ""
	slice nil
	map nil
	pointer nil
	channel nil
	interface nil
	function nil
	struct 字段零值
	*/

	p1 := Person{100.0}
	p2 := Person{10.0}
	if p1 == p2 {
		fmt.Println("equals")
	}

	// nil切片
	//var slc1 []Person
	// empty切片
	var slc2 = make([]Person, 0)
	if slc2 == nil {
		fmt.Println("nil slice")
	}

	// nil map
	var m1 map[string]string
	//m1["0"] = "zero"
	// empty map
	m2 := make(map[string]string, 0)
	m2["0"] = "zero"

	for key, value := range m1 {
		fmt.Println(key, value)
	}
	for key, value := range m2 {
		fmt.Println(key, value)
	}

}
