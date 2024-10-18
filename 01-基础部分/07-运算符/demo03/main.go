package main

import "fmt"

func main() {
	/*
		逻辑运算符
			&& : 逻辑与，两边都为真，才为真
			|| : 逻辑或，只要有一边为真，则为真
			! : 逻辑非，取反
	*/
	// 逻辑运算符
	a := true
	b := false

	// 逻辑与
	AndResult := a && b
	fmt.Printf("逻辑与: %t\n", AndResult)

	// 逻辑或
	OrResult := a || b
	fmt.Printf("逻辑或: %t\n", OrResult)

	// 逻辑非
	NotResult := !a
	fmt.Printf("逻辑非: %t\n", NotResult)
}
