package main

import "fmt"

func main() {
	// 数组是值类型
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1

	var arr3 [3]int
	arr3 = arr1
	// 打印arr1、arr2、arr3的内存地址和值
	fmt.Printf("arr1地址:%p,arr1值:%#v\n", &arr1, arr1)
	fmt.Printf("arr2地址:%p,arr2值:%#v\n", &arr2, arr2)
	fmt.Printf("arr3地址:%p,arr3值:%#v\n", &arr3, arr3)

	fmt.Println("=======================================")

	// 对arr2数组的第二个元素进行修改
	arr2[1] = 200
	// 对arr3数组的第三个元素进行修改
	arr3[2] = 300

	// 再打印arr1、arr2、arr3的内存地址和值
	fmt.Printf("arr1地址:%p,arr1值:%#v\n", &arr1, arr1)
	fmt.Printf("arr2地址:%p,arr2值:%#v\n", &arr2, arr2)
	fmt.Printf("arr3地址:%p,arr3值:%#v\n", &arr3, arr3)
}
