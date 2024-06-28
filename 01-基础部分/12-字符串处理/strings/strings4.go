package main

import (
	"fmt"
	"strings"
)

func main() {
	// 定义字符串
	str := "Hello World"
	// 检查子字符串是否存在
	contains1 := strings.Contains(str, "World") // true
	fmt.Println(contains1)
	// 检查子字符串是否存在
	contains2 := strings.Contains(str, "golang") // false
	fmt.Println(contains2)
}
