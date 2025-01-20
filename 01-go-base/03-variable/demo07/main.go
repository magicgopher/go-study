package main

import "fmt"

func main() {
	// Go语言变量注意事项6
	// 变量的零值。也叫默认值。
	var v1 int        // 默认值：0
	var v2 string     // 默认值：""
	var v3 bool       // 默认值：false
	var v4 float64    // 默认值：0.0
	var v5 complex128 // 默认值：(0+0i)
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(v5)
}
