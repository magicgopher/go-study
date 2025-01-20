package main

import (
	"fmt"
)

func main() {
	// make函数创建map
	m1 := make(map[int]string)
	fmt.Printf("m的类型:%T\n", m1)
	fmt.Printf("m是否为nil:%t\n", m1 == nil) // false

	m2 := make(map[string]int, 10)
	fmt.Printf("m2的类型:%T\n", m2)
	fmt.Printf("m2是否为nil:%t\n", m2 == nil) // false
}
