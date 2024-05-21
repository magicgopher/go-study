package main

import "fmt"

func main() {
	// 定义一个[4][3]的二维数组。（多维数组同理）
	var array [4][3]int
	array = [4][3]int{
		{1, 2, 3},    // [0,0] [0,1] [0,2]
		{4, 5, 6},    // [1,0] [1,1] [1,2]
		{7, 8, 9},    // [2,0] [2,1] [2,2]
		{10, 11, 12}, // [3,0] [3,1] [3,2]
	}
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Printf("arr[%d][%d]:%v\n", i, j, array[i][j])
		}
	}
}
