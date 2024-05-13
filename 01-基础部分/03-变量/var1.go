package main

import "fmt"

func main() {
	// Go语言变量声明的三种方式
	// 第一种：指定变量数据类型，然后再赋值
	// 第二种：类型推断
	// 第三种：简短定义
	// 第四种：同时声明多个变量

	// 第一种：指定变量数据类型，然后再赋值
	// 写在一行
	var a int = 100
	fmt.Println("a:", a)

	// 分开写
	var b int
	b = 200
	fmt.Println("b:", b)

	// 第二种：类型推断
	var c = 300
	fmt.Println("c:", c)

	// 第三种：简短定义
	d := 400
	fmt.Println("d:", d)

	// 第四种：同时声明多个变量
	var e, f, g int = 500, 600, 700
	fmt.Println("e:", e)
	fmt.Println("f:", f)
	fmt.Println("g:", g)
}
