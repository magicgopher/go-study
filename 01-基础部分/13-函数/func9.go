package main

import "fmt"

func main() {
	// 匿名函数顾名思义就是没有名字的函数
	// 函数调用
	f1()

	// 匿名函数定义
	func() {
		fmt.Println("这是一个匿名函数")
	}()

	// 使用匿名函数定义一个加法函数
	res1 := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println(res1)
}

// 正常的函数
func f1() {
	fmt.Println("我是f1()函数...")
}
