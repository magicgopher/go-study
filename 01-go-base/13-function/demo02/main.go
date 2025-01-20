package main

import "fmt"

func main() {
	// 定义两个变量
	p1 := 20
	p2 := 30
	// 调用函数传入的p1,p2就是实际参数
	// 传入的p1就是函数的参数a,传入的p2就是函数的参数b
	// 注意事项：实参与形参必须一一对应：顺序，个数，类型
	res := add(p1, p2)
	fmt.Println(res)
}

// 定义一个函数用于计算两个 int 类型的数字之和
// 定义函数时 a,b 两个就是形式参数
func add(a, b int) int {
	return a + b
}
