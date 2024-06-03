package main

import "fmt"

func main() {
	// 定义一个三维切片
	threeDimSlice := make([][][]int, 2)

	// 为每个二维切片分配内存空间
	for i := range threeDimSlice {
		threeDimSlice[i] = make([][]int, 3)
		for j := range threeDimSlice[i] {
			threeDimSlice[i][j] = make([]int, 4)
		}
	}

	// 填充数据
	threeDimSlice[0][0] = []int{1, 2, 3, 4}
	threeDimSlice[0][1] = []int{5, 6, 7, 8}
	threeDimSlice[0][2] = []int{9, 10, 11, 12}
	threeDimSlice[1][0] = []int{13, 14, 15, 16}
	threeDimSlice[1][1] = []int{17, 18, 19, 20}
	threeDimSlice[1][2] = []int{21, 22, 23, 24}

	// 访问元素
	fmt.Println(threeDimSlice[1][1][2]) // 输出: 19
}
