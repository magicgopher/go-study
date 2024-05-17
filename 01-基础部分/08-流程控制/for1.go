package main

import "fmt"

func main() {
	/*
		for 循环语句基本形式
		语法格式：
			for 条件表达式 {
				// 循环体
			}
		注意事项：条件表达式只有true和false两种结果，true表示循环体执行，false表示循环体不执行
	*/
	// 循环语句 基本形式示例
	// 求1-100的和
	sum := 0
	i := 1
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println("1-100的和:", sum)
}
