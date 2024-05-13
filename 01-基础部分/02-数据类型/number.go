package main

import "fmt"

func main() {
	// 数值类型
	// 无符号整数
	var n1 uint8 = 255
	var n2 uint16 = 65535
	var n3 uint32 = 4294967295
	var n4 uint64 = 18446744073709551615

	// 有符号整数
	var n5 int8 = -128
	var n6 int16 = -32768
	var n7 int32 = -2147483648
	var n8 int64 = -9223372036854775808

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println("============")
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}
