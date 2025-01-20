package main

import "fmt"

func main() {
	/*
		高阶函数：
			根据go语言的数据类型的特点，可以将一个函数作为另一个函数的参数。

		fun1(),fun2()
		将fun1函数作为fun2函数的参数
			fun2函数：被称为高阶函数
				接收了一个函数作为参数的函数，高阶函数
			fun1函数：被称为回调函数
	*/

	// 设计一个函数，用于两个数的加减乘除运算符
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", operate)

	// 这里使用函数名()的方式调用函数，再将调用函数获得的结果赋值给res1
	res1 := add(1, 2)
	fmt.Println(res1)

	// 使用operate函数 加法
	res2 := operate(10, 2, add)
	fmt.Println(res2)

	// 减法
	res3 := operate(10, 1, sub)
	fmt.Println(res3)

	// 乘法
	res4 := operate(5, 10, mul)
	fmt.Println(res4)

	// 除法
	res5 := operate(90, 3, div)
	fmt.Println(res5)
}

// 加法
func add(a, b int) int {
	return a + b
}

// 减法
func sub(a, b int) int {
	return a - b
}

// 乘法
func mul(a, b int) int {
	return a * b
}

// 除法
func div(a, b int) int {
	return a / b
}

// 加减乘除操作
func operate(a, b int, fun func(int, int) int) int {
	// fmt.Println(a, b, fun) // 打印三个参数信息
	res := fun(a, b)
	return res
}
