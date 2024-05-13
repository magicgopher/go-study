package main

import "fmt"

func main() {
	// Go语言变量注意事项4
	// 简短定义方式，左边的变量名至少有一个是新的。
	age := 26               // 简短声明，声明age变量并赋值为26
	age, name := 30, "John" // 错误，age已经声明过

	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
}
