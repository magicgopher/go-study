package main

import "fmt"

func main() {
	// 定义切片长度
	const elementCount = 10
	// 定义原始切片
	srcData := make([]int, elementCount)
	// 给原始切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用原始切片
	refData := srcData
	// 定义拷贝切片
	copyData := make([]int, elementCount)
	// 拷贝原始切片
	copy(copyData, srcData)
	// 修改原始切片第一个元素
	srcData[0] = 999
	// 打印引用切片第一个元素
	fmt.Println(refData[0])
	// 打印拷贝切片第一个元素和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	copy(copyData, srcData[4:6])
	fmt.Println(copyData)
}
