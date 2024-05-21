package main

import "fmt"

func main() {
	// 定义一个长度为5的字符串数组
	var str1 [5]string = [5]string{"C++", "Go", "Rust", "Java", "Python"}

	// 使用for range遍历数组
	for index, value := range str1 {
		println(index, value)
	}

	// 通过数组元素索引修改数组元素
	str1[0] = "C"

	// 再次遍历数组元素
	fmt.Println("==================")
	for index, value := range str1 {
		println(index, value)
	}
}
