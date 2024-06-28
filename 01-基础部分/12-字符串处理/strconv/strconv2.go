package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串与布尔类型之间的转换
	// ParseBool函数的作用：将字符串转换为布尔类型
	b1, _ := strconv.ParseBool("true")
	fmt.Printf("b1类型:%T,b1值:%v\n", b1, b1)
	// Bool函数的作用：将布尔类型转换为字符串
	b2 := strconv.FormatBool(true)
	fmt.Printf("b2类型:%T,b2值:%v\n", b2, b2)
}
