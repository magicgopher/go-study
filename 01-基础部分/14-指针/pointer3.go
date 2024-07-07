package main

import "fmt"

func main() {
	// 通过 * 获取指针指向的变量的值
	s1 := "hello"
	p1 := &s1
	fmt.Printf("p1指针指向的变量的值是:%s\n", *p1)
	// 修改指针变量的值
	*p1 = "world"
	fmt.Printf("s1的值是:%s\n", s1)
}
