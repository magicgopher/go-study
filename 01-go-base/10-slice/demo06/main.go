package main

import "fmt"

func main() {
	// make函数创建string类型切片
	// 使用 make 创建的切片都不可能是 nil
	s1 := make([]string, 0)
	fmt.Println(len(s1)) // 0
	// 判断make函数创建的切片是否为nil
	fmt.Println(s1 == nil)
}
