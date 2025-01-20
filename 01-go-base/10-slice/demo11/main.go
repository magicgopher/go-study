package main

import "fmt"

func main() {
	// 直接移动数据指针的方式删除切片中的元素
	// 删除开头N个元素语法格式：切片名 = 切片名[N:]
	// 定义切片初始化5个元素
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1) // [1 2 3 4 5]
	s1 = s1[1:]     // 删除第一个元素【从下标1开始截取到末尾】
	fmt.Println(s1) // [2 3 4 5]
}
