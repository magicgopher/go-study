package main

import "fmt"

func main() {
	// 创建map
	m1 := map[string]string{
		"name":    "MagicGopher",
		"age":     "18",
		"address": "银河系地球",
	}
	// map可以通过len函数获取map的长度
	fmt.Printf("m1的长度:%d\n", len(m1))
	// map是无法通过cap函数获取map的容量的
	// fmt.Printf("m1的容量:%d\n", cap(m1))
	// map的值
	fmt.Printf("m1的值:%v\n", m1)
}
