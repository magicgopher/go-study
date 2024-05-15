package main

import "fmt"

// 定义类型
type ByteSize float64

func main() {
	// 定义数量级
	const (
		_           = iota             // 通过分配给空白标识符来忽略第一个值
		KB ByteSize = 1 << (10 * iota) // 等价于 1 << (10 * 1)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	// 打印输出
	fmt.Printf("%0.1f\n", KB)
	fmt.Printf("%0.1f\n", MB)
	fmt.Printf("%0.1f\n", GB)
	fmt.Printf("%0.1f\n", TB)
	fmt.Printf("%0.1f\n", PB)
	fmt.Printf("%0.1f\n", EB)
	fmt.Printf("%0.1f\n", ZB)
	fmt.Printf("%0.1f\n", YB)
}
