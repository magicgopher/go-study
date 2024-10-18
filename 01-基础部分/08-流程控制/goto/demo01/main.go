package main

import "fmt"

func main() {
	// goto: 提供了一种无条件跳转的控制流方式
	num := 100
	fmt.Println("num:", num)
	goto Label1 // 跳转到Label1标签处
	num = 200   // 这行代码不会执行
Label1:
	fmt.Println("num:", num)
}
