package main

import "fmt"

func main() {
	// 打印一下func1函数的数据类型
	fmt.Printf("%T\n", func1) // 结果：func(int, int)
	fmt.Println(func1) // 这里输出的是函数的内存地址：0x104436db0

	// 定义一个函数类型的变量
	var f1 func(int, int)
	fmt.Println(f1) // <nil> 空指针未分配状态

	// 给 f1 函数变量赋值，这里必须保证赋值的数据和声明的数据类型是一致的
	f1 = func1 // 将func1的值(函数体的地址)赋值给f1
	fmt.Println(f1) // 这里f1输出的就是func1的地址：0x104436db0

	// 调用func1函数
	func1(10, 20)
	// 调用f1函数变量
	f1(11, 22)
}

// 定义一个fun1函数
func func1(a, b int) {
	fmt.Printf("a:%d, b:%d\n", a, b)
}