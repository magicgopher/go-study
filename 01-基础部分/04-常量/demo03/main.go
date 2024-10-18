package main

import "fmt"

func main() {
	/*
		iota 是 Go 语言提供的一个关键字，它表示自增的整数常量。
		iota 的值从 0 开始，每次递增 1。
		语法格式：
			const (
				枚举1 = iota
				枚举2
				枚举3
				...
				枚举n
			)
	*/

	// 口诀：iota，自动递增，每行加一，不用猜。
	const (
		a = iota    // 这里iota的值是0
		b = "hello" // b的值是hello iota=1
		c           // c的值默认和上一行一致是hello iota=2
		d = 6.66    // d的值是6.66 iota=3
		e = "world" // e的值是world iota=4
		f = true    // f的值是true iota=5
		g = iota    // g的值是6 iota=6
		h           // h的值是7 iota=7
	)
	fmt.Println(a, b, c, d, e, f, g, h)

	// iota还可以与位操作符结合使用，用于定义按位运算的常量。
	const (
		FlagA = 1 << iota // 1
		FlagB             // 2
		FlagC             // 4
	)
	fmt.Println(FlagA, FlagB, FlagC)
}
