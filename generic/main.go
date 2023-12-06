package main

import "fmt"

// 类型集
type TypeSet interface {
	// ~T代表底层类型是T的所有类型
	int | int64 | string | ~float64
}

type f64 float64

func max[T TypeSet](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 泛型函数
func log[T any](slc []T) {
	for _, value := range slc {
		fmt.Println(value)
	}
}

// 泛型结构体
type GS[T int | string] struct {
	id T
}

// 泛型接口
type GI[T any] interface {
	log(slc []T)
}

// 泛型切片
type gs[T int | float32 | float64] []T

// 泛型方法:绑定到某个自定义泛型类型
func (gs gs[T]) sum() T {
	var sum T
	for _, value := range gs {
		sum += value
	}
	return sum
}

func main() {
	// 泛型的使用场景:自定义泛型类型,泛型方法,泛型函数
	slc1 := []string{"0", "5"}
	slc2 := []float64{0.0, 5.0}
	log(slc1)
	log(slc2)

	var gs1 gs[int] = []int{0, 5}
	fmt.Println(gs1.sum())
	fmt.Printf("%T\n", gs1)
	var gs2 gs[float32] = []float32{0.0, 5.0}
	fmt.Printf("%T\n", gs2)

	// 泛型map
	// 类型参数K,V跟着的是类型约束
	type gm[K string, V any] map[K]V
	var gm1 gm[string, struct{}] = map[string]struct{}{
		"0": {},
	}
	fmt.Println(gm1)

	// 泛型通道
	type gc[T int | string] chan T
	var gc1 gc[string] = make(chan string, 5)
	gc1 <- "0"

	// 泛型指针
	type gp[T any] *T
	var gp1 gp[interface{}] = new(interface{})
	fmt.Println(gp1)

	var a f64 = 0.0
	var b f64 = 5.0
	fmt.Println(max(a, b))
}
