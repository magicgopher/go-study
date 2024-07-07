package main

import "fmt"

func main() {
	// 函数：是一个独立的代码块，用于执行特定的任务或计算。
	fmt.Println("main函数开始执行")
	// fmt.Println("Hello")
	// fmt.Println("World")
	// 将打印输出Hello和打印输出World抽离成一个函数，再调用
	printInfo()
	fmt.Println("main函数结束执行")
}

func printInfo() {
	fmt.Println("Hello")
	fmt.Println("World")
}
