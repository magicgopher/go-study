package main

import "fmt"

func main() {
	// 定义一个map
	var m1 map[int]string
	fmt.Printf("m的类型:%T\n", m1)
	fmt.Printf("m是否为nil:%t\n", m1 == nil) // true
}
