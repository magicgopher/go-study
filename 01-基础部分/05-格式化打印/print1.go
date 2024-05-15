package main

import "fmt"

func main() {
	/*
		格式化打印常用的函数，如下：
			// Print 输出到标准输出
			func Print(a ...any) (n int, err error)
			// Printf 格式化输出
			func Printf(format string, a ...any) (n int, err error)
			// Println 输出到标准输出，并换行
			func Println(a ...any) (n int, err error)
	*/

	// print函数的使用
	str := "Hello Go!!!"
	fmt.Print(str)
	num1 := 100
	fmt.Print(num1)
	f1 := 9.999
	fmt.Print(f1)
}
