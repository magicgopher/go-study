package main

import "fmt"

func main() {
	// 嵌套循环，先执行外层循环，再执行内层循环
	// 嵌套循环案例，打印99乘法表
	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d x %d = %d\t", y, x, x*y)
		}
		fmt.Println()
	}
}
