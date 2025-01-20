package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// 将 double 函数作为参数传递给 applyOperation 函数
	applyOperation(numbers, double)
}

// 定义一个接受函数作为参数的函数
func applyOperation(nums []int, operation func(int) int) {
	for _, num := range nums {
		result := operation(num)
		fmt.Printf("Operation result: %d\n", result)
	}
}

// 示例函数：加倍操作
func double(n int) int {
	return n * 2
}
