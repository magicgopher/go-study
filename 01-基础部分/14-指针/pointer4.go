package main

import "fmt"

func main() {
	// 使用new函数创建一个int类型的指针变量
	ptr := new(int)
	fmt.Printf("ptr的值是:%v\n", ptr)
	fmt.Printf("ptr指向的变量的值是:%v\n", *ptr)
	fmt.Printf("ptr自己的内存地址:%v\n", &ptr)
}
