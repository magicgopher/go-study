package main

import "fmt"

func main() {
	/*
		赋值运算符
			= : 赋值运算符
			+= : 相加后再赋值
			-= : 相减后再赋值
			*= : 相乘后再赋值
			/= : 相除后再赋值
			%= : 求余后再赋值
			<<= : 按位左移后再赋值
			>>= : 按位右移后再赋值
			&= : 按位与后再赋值
			^= : 按位异或后再赋值
			|= : 按位或后再赋值
			&^= : 按位非后再赋值
	*/
	// 赋值运算符
	var A = 10
	var B = 5
	var C int

	// 简单赋值运算符
	C = A + B
	fmt.Println("C =", C) // 输出：C = 15

	// 相加后再赋值
	C += A
	fmt.Println("C =", C) // 输出：C = 25

	// 相减后再赋值
	C -= A
	fmt.Println("C =", C) // 输出：C = 15

	// 相乘后再赋值
	C *= A
	fmt.Println("C =", C) // 输出：C = 150

	// 相除后再赋值
	C /= A
	fmt.Println("C =", C) // 输出：C = 15

	// 求余后再赋值
	C %= A
	fmt.Println("C =", C) // 输出：C = 5

	// 按位左移后赋值
	C <<= 2
	fmt.Println("C =", C) // 输出：C = 20

	// 按位右移后赋值
	C >>= 2
	fmt.Println("C =", C) // 输出：C = 5

	// 按位与后赋值
	C &= 2
	fmt.Println("C =", C) // 输出：C = 0

	// 按位异或后赋值
	C ^= 2
	fmt.Println("C =", C) // 输出：C = 2

	// 按位或后赋值
	C |= 2
	fmt.Println("C =", C) // 输出：C = 2

	// 按位非后赋值
	C &^= 2
	fmt.Println("C =", C) // 输出：C = 0
}
