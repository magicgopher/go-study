package main

import "fmt"

func main() {
	// 定义一个字符串
	s1 := "hello world!"
	// 获取子字符串 "hello"
	sub1 := s1[0:5] // 从索引0开始，到索引5结束，但不包含索引5
	fmt.Println(sub1)
	// 获取子字符串 "world!"
	// sub2 := s1[6:12] // 从索引6开始，到索引12结束，但不包含索引12
	// sub2 := s1[6:] // 从索引6开始，到字符串末尾结束
	sub2 := s1[6:len(s1)] // 从索引6开始，到字符串末尾结束，不包含最后的索引
	fmt.Println(sub2)
}
