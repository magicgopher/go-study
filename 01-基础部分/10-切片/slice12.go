package main

import "fmt"

func main() {
	// 定义切片
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1) // [1 2 3 4 5]
	// 通过 append 函数将第一个元素之后的所有元素复制到一个新的切片中，覆盖原切片内容
	s1 = append(s1[:0], s1[1:]...)
	fmt.Println(s1) // [2 3 4 5]
}
