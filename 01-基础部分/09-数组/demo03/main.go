package main

import "fmt"

func main() {
	// 如果两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下
	// 们可以直接通过较运算符（==和!=）来判断两个数组是否相等
	// 只有当两个数组的所有元素都是相等的时候数组才是相等的
	// 不能比较两个类型不同的数组，否则程序将无法完成编译

	// 定义两个数组，比较数组的元素是否相同
	var arr1 [3]int = [3]int{10, 20, 30}
	var arr2 [3]int = [3]int{10, 20, 30}
	fmt.Printf("arr1数组元素和arr2的元素是否相同:%t\n", arr1 == arr2)

	// var arr3 [2]int = [2]int{10, 20}
	// fmt.Printf("arr1数组元素和arr3的元素是否相同:%t\n", arr1 == arr3) // invalid operation: arr1 == arr3 (mismatched types [3]int and [2]int)
}
