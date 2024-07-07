package main

import "fmt"

func main() {
	// 定义一个int类型变量a
	a := 10
	fmt.Printf("a的值是%d,a的内存地址是:%p\n", a, &a) // a = 10
	// 定义一个指针变量p1，指向变量a
	p1 := &a
	fmt.Printf("p1的数据类型是:%T,p1存储的值是:%v,p1的内存地址是:%p\n", p1, p1, &p1)
	// 二级指针 这里存储的值是一级指针的内存地址 &p1
	p2 := &p1
	fmt.Printf("p2的数据类型是:%T,p2存储的值是:%v,p2的内存地址是:%p\n", p2, p2, &p2)
	// 三级指针 这里存储的值是二级指针的内存地址 &p2
	p3 := &p2
	fmt.Printf("p3的数据类型是:%T,p3存储的值是:%v,p3的内存地址是:%p\n", p3, p3, &p3)
	// 修改三级指针指向的变量存储的值
	***p3 = 111
	fmt.Printf("a的值是:%d\n", a) // a = 111
}
