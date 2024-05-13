package main

import "fmt"

func main() {
	// Go语言变量注意事项1
	// 变量必须先声明(定义)才能使用。
	var age int
	fmt.Println("Age:", age) // 正确，age已经声明

	// fmt.Println("Name:", name) // 错误，name未声明
	// name := "John"             // 错误，变量使用在声明之前

	var score int
	score = 90
	fmt.Println("Score:", score) // 正确，score已经声明并赋值
}
