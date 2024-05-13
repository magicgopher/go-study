package main

import "fmt"

func main() {
	// Go语言变量注意事项3
	// 变量名不能冲突。(同一个作用于域内不能冲突)
	var age int = 26
	fmt.Println("Age:", age)

	// var age int = 30  错误，重复声明同名变量

	var name string = "John"
	fmt.Println("Name:", name)

	// 简短变量声明 := 运算符用于在声明变量的同时进行赋值操作，它会自动推断变量的类型。
	// name := "Doe" // 错误，重复声明同名变量

	// ...
}
