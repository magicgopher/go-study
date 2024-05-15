package main

import "fmt"

func main() {
	/*
		同时声明多个常量。
			语法格式
			const (
				常量名1 [数据类型1] = 值1
				常量名2 [数据类型2] = 值2
				常量名3 [数据类型3] = 值3
				...
				常量名n [数据类型n] = 值n
			)
	*/

	// 声明多个常量
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
	fmt.Println(SPRING, SUMMER, AUTUMN, WINTER)

	// 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
	const (
		a = 100
		b         // 默认和上一行相同【数据类型和值都是一样的】
		c         // 默认和上一行相同【数据类型和值都是一样的】
		d float64 = 5.555
		e string  = "hello"
		f         // 默认和上一行相同【数据类型和值都是一样的】
		g rune    = 'A'
	)
	fmt.Println(a, b, c, d, e, f, g)
}
