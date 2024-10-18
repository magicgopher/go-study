package main

import "fmt"

func main() {
	// Go语言中常量的声明方式(常量名全部大写，多个单词需要使用_分开)
	// 显式类型定义：const 常量名 数据类型 = 值
	// 隐式类型定义：const 常量名 = 值

	// 显式类型定义
	// 指定了数据类型
	const PI float64 = 3.1415926
	fmt.Println(PI)

	// 隐式类型定义
	// 没有指定数据类型，根据编译器自动推断，这里的30会被推断为int类型
	const DAY = 30
	fmt.Println(DAY)
}
