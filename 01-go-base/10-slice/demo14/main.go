package main

import "fmt"

func main() {
	// 定义切片
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	// 移除后面一个元素
	s1 = s1[:len(s1)-1]
	fmt.Println(s1)
	// 移除后面N个元素
	s1 = s1[:len(s1)-2]
	fmt.Println(s1)
}
