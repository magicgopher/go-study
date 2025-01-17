package main

import "fmt"

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
	// 复数类型
	fmt.Println("复数类型")
	var c1 complex64 = 3.14159265358979 + 2.71828182845905i
	var c2 complex128 = 3.14159265358979323846264338327950288 + 2.71828182845904523536028747135266249i
	fmt.Println("c1 complex64类型变量的值是:", c1)
	fmt.Println("c2 complex128类型变量的值是:", c2)
}
