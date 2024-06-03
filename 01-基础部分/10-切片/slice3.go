package main

import "fmt"

func main() {
	// 定义一个int类型长度10的数组
	arr1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 打印数组信息
	fmt.Printf("len:%d\n", len(arr1)) // len:10
	fmt.Printf("arr1:%T\n", arr1)     // arr1:[10]int

	// 将数组索引为2到5的元素转成切片
	// 这里2到5是左闭右开的等价于[2,5)表示一个开区间，其中包含2但不包含5
	// 长度等于结束索引减去起始索引，所以长度为3
	s1 := arr1[2:5]
	fmt.Printf("len:%d\n", len(s1)) // len:3
	fmt.Printf("s1:%T\n", s1)       // s1:[]int
}
