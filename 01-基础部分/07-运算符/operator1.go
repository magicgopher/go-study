package main

import "fmt"

func main() {
	/*
		算数运算符
		+ : 加法
		- : 减法
		* : 乘法
		/ : 除法
		% : 取模
		++ : 自增（只能独占一行，不能参与其他运算）
		-- : 自减（只能独占一行，不能参与其他运算）
	*/

	// 算数运算符示例
	a := 10
	b := 20

	// 加
	sum := a + b
	fmt.Printf("和: %d\n", sum)

	// 减
	diff := a - b
	fmt.Printf("差: %d\n", diff)

	// 乘
	prod := a * b
	fmt.Printf("积: %d\n", prod)

	// 除
	quot := b / a
	fmt.Printf("商: %d\n", quot)

	// 取模
	rem := b % a
	fmt.Printf("余数: %d\n", rem)

	// 自增
	a++
	fmt.Printf("a: %d\n", a)

	// 自减
	b--
	fmt.Printf("b: %d\n", b)
}
