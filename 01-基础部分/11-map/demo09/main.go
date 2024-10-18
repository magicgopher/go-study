package main

import "fmt"

func main() {
	// 创建一个 map
	m := map[string]int{
		"苹果": 5,
		"香蕉": 3,
		"橙子": 7,
	}

	// 检查 "香蕉" 键是否存在
	value, ok := m["香蕉"]
	if ok {
		fmt.Printf("'香蕉' 的值是: %d\n", value)
	} else {
		fmt.Println("'香蕉' 键不存在于 map 中")
	}

	// 检查 "梨" 键是否存在
	value, ok = m["梨"]
	if ok {
		fmt.Printf("'梨' 的值是: %d\n", value)
	} else {
		fmt.Println("'梨' 键不存在于 map 中")
	}
}
