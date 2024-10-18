package main

import "fmt"

func main() {
	// 表达式 switch语句格式
	// switch语句示例，多个case放在一行示例
	condition := 5
	switch condition {
	case 1, 3, 5, 7, 9:
		fmt.Println(true)
	case 2, 4, 6, 8, 10:
		fmt.Println(false)
	default:
		fmt.Println(0)
	}
}
