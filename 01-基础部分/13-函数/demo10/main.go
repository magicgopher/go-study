package main

import "fmt"

func main() {
	// 调用f1函数
	f1()
	// 匿名函数，在main函数中只会执行一次，但是使用函数变量的方式可以执行多次
	// 定义一个函数变量
	var f2 func()
	f2 = func() {
		fmt.Println("这是一个匿名函数...")
	}
	// 打印数据类型
	fmt.Printf("f2:%T\n", f2)
	// 调用匿名函数
	f2()
	f2()

	// 定义一个有返回类型的匿名函数
	res1 := func(a, b int) int {
		return a + b
	} // 这里并没有传入参数，所以就是将匿名函数赋值给res1变量
	fmt.Printf("res1数据类型:%T\n", res1)
	res2 := func(a, b int) int {
		return a + b
	}(5, 6) // 这里传入了参数进行了函数的调用，会将结果的数据赋值给res2
	fmt.Printf("res2数据类型:%T\n", res2)
}

// 正常的函数
func f1() {
	fmt.Println("这是f1()函数...")
}
