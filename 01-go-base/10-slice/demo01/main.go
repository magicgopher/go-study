package main

import "fmt"

func main() {
	// 定义一个int类型的切片
	var s1 []int
	fmt.Printf("s1的类型:%T\n", s1) // s1的类型:[]int

	// 使用make函数定义一个string类型的切片
	s2 := make([]string, 0, 3)
	fmt.Printf("s2的类型:%T\n", s2) // s2的类型:[]string
}
