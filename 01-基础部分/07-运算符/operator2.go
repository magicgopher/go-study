package main

import "fmt"

func main() {
	/*
		关系运算符
			== : 判断两个操作数是否相等
			!= : 判断两个操作数是否不相等
			> : 判断第一个操作数是否大于第二个操作数
			< : 判断第一个操作数是否小于第二个操作数
			>= : 判断第一个操作数是否大于等于第二个操作数
			<= : 判断第一个操作数是否小于等于第二个操作数
	*/
	// 关系运算符示例
	a := 10
	b := 20

	// ==
	equal := a == b
	fmt.Printf("相等: %t\n", equal)

	// !=
	notEqual := a != b
	fmt.Printf("不相等: %t\n", notEqual)

	// >
	greaterThan := a > b
	fmt.Printf("大于: %t\n", greaterThan)

	// <
	lessThan := a < b
	fmt.Printf("小于: %t\n", lessThan)

	// >=
	greaterOrEqual := a >= b
	fmt.Printf("大于等于: %t\n", greaterOrEqual)

	// <=
	lessOrEqual := a <= b
	fmt.Printf("小于等于: %t\n", lessOrEqual)
}
