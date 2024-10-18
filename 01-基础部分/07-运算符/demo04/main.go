package main

import "fmt"

func main() {
	/*
		位运算符
			& : 按位与运算符
			| : 按位或运算符
			^ : 按位异或运算符
			&^: 按位清除运算符
			<<: 按位左移运算符
			>>: 按位右移运算符
	*/
	// 位运算符示例
	a := 60 // 0011 1100
	b := 13 // 0000 1101

	// 按位与运算符
	And := a & b
	fmt.Printf("按位与: %d (%08b)\n", And, And)

	// 按位或运算符
	Or := a | b
	fmt.Printf("按位或: %d (%08b)\n", Or, Or)

	// 按位异或运算符
	Xor := a ^ b
	fmt.Printf("按位异或: %d (%08b)\n", Xor, Xor)

	// 按位清除运算符
	AndNot := a &^ b
	fmt.Printf("按位清除: %d (%08b)\n", AndNot, AndNot)

	// 按位左移运算符
	Left := a << 2
	fmt.Printf("按位左移: %d (%08b)\n", Left, Left)

	// 按位右移运算符
	Right := a >> 2
	fmt.Printf("按位右移: %d (%08b)\n", Right, Right)
}
