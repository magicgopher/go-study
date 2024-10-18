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
	// 字节类型
	fmt.Println("字节类型")
	var b byte = 66 // 表示字符 B
	fmt.Printf("byte类型变量b的值是:%d,对应ASCII字符串:%c\n", b, b)
	fmt.Println("byte类型等同于uint8")
	var b1 uint8 = b + 1 // 66 + 1 = 67 对应 ASCII 编码表为 C
	fmt.Printf("b1 uint8类型变量的值是:%d,对应ASCII字符串:%c\n", b1, b1)
	fmt.Println("============")
	var r rune = 'G'
	fmt.Printf("rune类型变量r的值是:%d,对应ASCII字符:%c\n", r, r)
	fmt.Println("rune类型等同于int32类型")
	var r1 int32 = r + 1
	fmt.Printf("rune类型变量r1的值是:%d,对应ASCII字符:%c\n", r1, r1)
	// Unicode字符
	var r2 rune = '你'
	fmt.Printf("rune类型变量r2的值是%d,对应Unicode字符是:%c\n", r2, r2)
}
