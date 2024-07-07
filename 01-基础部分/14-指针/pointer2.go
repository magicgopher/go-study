package main

import "fmt"

func main() {
	// 通过 & 获取变量的内存地址
	s1 := "hello"
	fmt.Printf("s1的地址值:%p\n", &s1)
}
