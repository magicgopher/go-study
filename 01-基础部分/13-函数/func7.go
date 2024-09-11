package main

import "fmt"

// 返回一个函数的函数
func multiplier(factor int) func(int) int {
	// 返回一个闭包函数
	return func(n int) int {
		return n * factor
	}
}

func main() {
	// 创建一个乘以 3 的函数
	times3 := multiplier(3)

	// 使用这个函数来乘以不同的数值
	result1 := times3(5) // 5 * 3 = 15
	result2 := times3(10) // 10 * 3 = 30

	fmt.Println("Result 1:", result1)
	fmt.Println("Result 2:", result2)
}