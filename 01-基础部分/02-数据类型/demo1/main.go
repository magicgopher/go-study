package main

import (
	"fmt"
)

func main() {
	/*
		go语言的数据类型分成两类：一类是基本类型、另一类是派生类型（复合类型）
		基本类型有数值类型、浮点类型、布尔类型、复数类型、字符串类型、字节类型
			数值类型：
				有符号位：int、int8、int16、int32、int64
				无符号位：uint、uint8、uint16、uint32、uint64（无符号位是不可以存储负数的）
			浮点类型：float32、float64
			复数类型：complex64、complex128
			布尔类型：bool
			字符串类型：string
			字节类型：byte、rune，其中byte等同于uint8，rune等同于int32，用于表示UTF-8字符串的Unicode码点
	*/
	// 数值类型
	fmt.Println("数值类型-无符号位")
	// 数值类型
	// 无符号整数
	var n1 uint = 100
	var n2 uint8 = 255
	var n3 uint16 = 65535
	var n4 uint32 = 4294967295
	var n5 uint64 = 18446744073709551615
	fmt.Println("uint类型变量n1的值是:", n1)
	fmt.Println("uint8类型变量n2的值是:", n2)
	fmt.Println("uint16类型变量n3的值是:", n3)
	fmt.Println("uint32类型变量n4的值是:", n4)
	fmt.Println("uint64类型变量n5的值是:", n5)
	fmt.Println("============")
	// 有符号整数
	fmt.Println("数值类型-无符号位")
	var n6 int = -100
	var n7 int8 = -128
	var n8 int16 = -32768
	var n9 int32 = -2147483648
	var n10 int64 = -9223372036854775808
	fmt.Println("int类型变量n6的值是:", n6)
	fmt.Println("int8类型变量n7的值是:", n7)
	fmt.Println("int16类型变量n8的值是:", n8)
	fmt.Println("int32类型变量n9的值是:", n9)
	fmt.Println("int64类型变量n10的值是:", n10)
}
