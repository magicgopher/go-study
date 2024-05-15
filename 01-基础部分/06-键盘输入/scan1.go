package main

import "fmt"

func main() {
	/*
		使用fmt提供的键盘读取函数获取键盘输入的内容
			常用的几个函数如下：
				// Scan 函数从标准输入读取空格分隔的值,将数据保存到传入的参数中。
				// 它返回成功读取的项数和遇到的任何错误。
				func Scan(a ...any) (n int, err error)

				// Scanf 函数从标准输入读取数据,根据格式字符串格式化输入并将结果保存到传入的参数中。
				// 它返回成功读取的项数和遇到的任何错误。
				func Scanf(format string, a ...any) (n int, err error)

				// Scanln 函数从标准输入读取数据,直到遇到换行符为止,将数据保存到传入的参数中。
				// 它返回成功读取的项数和遇到的任何错误。
				func Scanln(a ...any) (n int, err error)
	*/

	var name string
	var age int

	// Scan函数
	fmt.Println("请输入姓名：")
	fmt.Scan(&name)
	fmt.Println("请输入年龄：")
	fmt.Scan(&age)

	// 格式化打印
	fmt.Printf("姓名:%s, 年龄:%d\n", name, age)
}
