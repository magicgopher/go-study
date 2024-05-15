package main

import "fmt"

func main() {
	// 运算符优先级示例
	a := 5
	b := 3
	c := 2

	// 表达式计算顺序
	result1 := a + b*c // 先计算 b * c，再加 a
	// result1 = 5 + 3 * 2 = 11

	result2 := (a + b) * c // 先计算 a + b，再乘以 c
	// result2 = (5 + 3) * 2 = 16

	result3 := a * b / c // 先计算 a * b，再除以 c
	// result3 = 5 * 3 / 2 = 7

	result4 := a%b + c // 先计算 a % b，再加 c
	// result4 = 5 % 3 + 2 = 4

	result5 := a < b && b < c // 先计算 a < b，再计算 b < c，最后进行逻辑与运算
	// result5 = false && true = false

	fmt.Println("result1:", result1)
	fmt.Println("result2:", result2)
	fmt.Println("result3:", result3)
	fmt.Println("result4:", result4)
	fmt.Println("result5:", result5)
}
