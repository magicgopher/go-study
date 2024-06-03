package main

import "fmt"

func main() {
	// 使用append函数给切片末尾追加元素。
	// append函数只能在末尾追加元素，不能在中间【或者指定索引位置】插入元素。
	// 注意append函数会返回一个新切片，所以这里需要重新赋值给s1
	s1 := make([]int, 5, 10) // 这里初始化的切片前面5个元素是0，后面5个元素是nil
	fmt.Printf("s1类型:%T\n", s1)
	fmt.Printf("s1长度:%d,s1容量:%d\n", len(s1), cap(s1)) // len: 5, cap: 10
	fmt.Printf("s1:%v\n", s1)
	fmt.Println("=======================")
	fmt.Println("append()第一次")
	s1 = append(s1, 1, 2, 3, 4, 5) // 5个元素
	fmt.Printf("s1类型:%T\n", s1)
	fmt.Printf("s1长度:%d,s1容量:%d\n", len(s1), cap(s1))
	fmt.Printf("s1:%v\n", s1)
	fmt.Println("=======================")
	fmt.Println("append()第二次") // 第二次追加元素，发现容量不够，自动扩容
	s1 = append(s1, 6, 7, 8)   // 3个元素
	fmt.Printf("s1类型:%T\n", s1)
	fmt.Printf("s1长度:%d,s1容量:%d\n", len(s1), cap(s1)) // len: 13, cap: 20
	fmt.Printf("s1:%v\n", s1)
}
