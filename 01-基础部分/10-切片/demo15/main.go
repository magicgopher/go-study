package main

import "fmt"

func main() {
	// 定义一个二维切片
	twoDimSlice := make([][]int, 3)

	// 为每个一维切片分配内存空间
	for i := range twoDimSlice {
		twoDimSlice[i] = make([]int, 4)
	}

	// 填充数据
	twoDimSlice[0] = []int{1, 2, 3, 4}
	twoDimSlice[1] = []int{5, 6, 7, 8}
	twoDimSlice[2] = []int{9, 10, 11, 12}

	// 访问元素
	fmt.Println(twoDimSlice[1][2]) // 输出: 7
}
