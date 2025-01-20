package main

import "fmt"

func main() {
	// 定义string类型切片
	// 此时的切片为空【栈内存已经分配好切片的指针、长度、容量】
	// 但是堆内存中还没有给切片分配任何存储空间
	var s1 []string // nil
	// 定义int类型切片
	// 这里原理同上【和s1切片同理】
	var s2 []int // nil
	// 定义一个int类型空切片
	s3 := []int{} // {}
	// 打印三个切片
	fmt.Println(s1, s2, s3) // [] [] []
	// 打印三个切片的长度
	fmt.Println(len(s1), len(s2), len(s3)) // 0 0 0
	// 切片判断为空
	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) // true
	fmt.Println(s3 == nil) // false
}
