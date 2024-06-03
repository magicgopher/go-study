package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8}
	// 拷贝之前
	fmt.Printf("s1:%v\n", s1) // [1 2 3 4 5]
	fmt.Printf("s1的长度:%d,s1的容量:%d\n", len(s1), cap(s1))
	fmt.Printf("s2:%v\n", s2) // [6 7 8]
	fmt.Printf("s2的长度:%d,s2的容量:%d\n", len(s2), cap(s2))
	// 将s1拷贝到s2
	copy(s2, s1)
	fmt.Printf("s1:%v\n", s1) // [1 2 3 4 5]
	fmt.Printf("s1的长度:%d,s1的容量:%d\n", len(s1), cap(s1))
	fmt.Printf("s2:%v\n", s2) // [1 2 3]
	fmt.Printf("s2的长度:%d,s2的容量:%d\n", len(s2), cap(s2))
}
