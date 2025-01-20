package main

import "fmt"

func main() {
	// 定义一个切片int类型
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(s1)
	// 移除值为10的元素
	s1 = append(s1[:9], s1[10:]...)
	fmt.Println(s1)
}
