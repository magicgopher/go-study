package main

import "fmt"

func main() {
	// 将数组转成切片语法格式：数组名[起始索引:结束索引]
	// 创建一个int类型长度10的数组
	arr1 := [10]int{27, 86, 82, 48, 80, 83, 19, 68, 96, 79}
	fmt.Printf("arr1的类型:%T\n", arr1)
	fmt.Printf("arr1的长度:%d,容量:%d\n", len(arr1), cap(arr1))
	fmt.Printf("arr1的内容:%v\n", arr1)
	fmt.Println("=======================================")

	// 将数组转成切片,这里的[:]表示将数组的全部元素拷贝到切片中
	s1 := arr1[:]
	fmt.Printf("s1的类型:%T\n", s1)
	fmt.Printf("s1的长度:%d,容量:%d\n", len(s1), cap(s1))
	fmt.Printf("s1的内容:%v\n", s1)
	fmt.Println("=======================================")

	// 从:左边没有写就是从0开始拷贝元素,直到拷贝到:右边的索引为4位置,索引为5是不包含的.
	s2 := arr1[:5]
	fmt.Printf("s2的类型:%T\n", s2)
	fmt.Printf("s2的长度:%d,容量:%d\n", len(s2), cap(s2))
	fmt.Printf("s2的内容:%v\n", s2)
	fmt.Println("=======================================")

	// 从:右边没有写就是从索引为6位置开始拷贝元素,直到拷贝到数组的最后一个元素.
	s3 := arr1[6:]
	fmt.Printf("s3的类型:%T\n", s3)
	fmt.Printf("s3的长度:%d,容量:%d\n", len(s3), cap(s3))
	fmt.Printf("s3的内容:%v\n", s3)
}
