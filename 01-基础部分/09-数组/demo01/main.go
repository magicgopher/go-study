package main

import "fmt"

func main() {
	// 数组：固定长度的集合，每个元素都是相同类型的变量

	// 定义一个int类型的数组
	var arr1 [5]int = [5]int{33, 28, 37, 45, 95}

	// 通过数组的索引访问数组的元素，索引从0开始，语法格式：数组名[索引]
	fmt.Println("arr1[1]的内容:", arr1[1]) // 打印数组索引下标为1的元素内容

	// 可以通过len()内置函数来获取数组的长度
	fmt.Println("数组的长度:", len(arr1))

	// 通过for循环遍历数组
	for i := 0; i < len(arr1); i++ {
		fmt.Println("arr1[", i, "]的内容:", arr1[i])
	}
}
