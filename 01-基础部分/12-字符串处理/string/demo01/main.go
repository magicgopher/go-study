package main

import "fmt"

func main() {
	// 定义单行字符串语法格式
	str1 := "Hello, World!"
	// 定义多行字符串语法格式
	str2 := `
	Hello, World!
	你好, 世界!
	`
	// // Hello, World!
	fmt.Println(str1)
	// Hello, World! 你好, 世界! 分成两行显示
	fmt.Println(str2)
}
