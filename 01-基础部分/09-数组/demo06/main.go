package main

import "fmt"

func main() {
	// 定义一个int类型的数组，长度为3
	arr1 := [3]int{1, 2, 3}
	fmt.Printf("arr1=%#v\n", arr1)

	// arr1 = [4]int{1, 2, 3, 4} // [3]int 和 [4]int 是两种不同的数组类型
	arr1 = [3]int{11, 22, 33}
	fmt.Printf("arr1=%#v\n", arr1)
}
