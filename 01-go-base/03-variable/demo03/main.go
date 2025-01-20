package main

import "fmt"

func main() {
	// Go语言变量注意事项2
	// Go语言是静态语言，要求变量的类型和赋值的类型必须一致。
	var age int // 声明一个整数类型的变量age
	age = 26    // 赋值整数值给age

	var name string // 声明一个字符串类型的变量name
	name = "John"   // 赋值字符串值给name

	var ratio float64 // 声明一个浮点数类型的变量ratio
	ratio = 3.14      // 赋值浮点数值给ratio

	// var score int = 90 也可以在声明时赋值

	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Ratio:", ratio)
}
