package main

import "fmt"

func main() {
	printType("Hello Go!")
	printType(66)
	printType(1.024)
	printType(true)
	printType([]int{1, 1, 2}) // 未知类型
}

// 打印类型 参数是一个空接口
func printType(i interface{}) {
	switch v := i.(type) {
	case string: // 字符串类型执行
		fmt.Println("String value:", v)
	case int: // int类型执行
		fmt.Println("Integer value:", v)
	case float64: // float64类型执行
		fmt.Println("Float value:", v)
	case bool: // bool类型执行
		fmt.Println("Bool type:", v)
	default:
		fmt.Println("Unknown type")
	}
}
