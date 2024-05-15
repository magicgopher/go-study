package main

import "fmt"

func main() {
	// 接下来是Printf函数格式化打印的整数类型占位符示例
	var a = 67
	fmt.Printf("a:%b\n", a) // a:1000011 二进制
	fmt.Printf("a:%c\n", a) // a:C ASCII码对应的字符
	fmt.Printf("a:%d\n", a) // a:67 十进制
	fmt.Printf("a:%o\n", a) // a:103 八进制
	fmt.Printf("a:%O\n", a) // a:0o103 八进制 前缀0o
	fmt.Printf("a:%q\n", a) // a:'C' 单引号包裹的字符
	fmt.Printf("a:%x\n", a) // a:43 十六进制 小写
	fmt.Printf("a:%X\n", a) // a:43 十六进制 大写
	fmt.Printf("a:%U\n", a) // a:U+0043 以 Unicode 格式打印整数值

	fmt.Println("=======================")

	// 接下来是浮点数的格式化打印的占位符示例
	f1 := 6.123456
	fmt.Printf("f1:%f\n", f1)    // f1:6.123456 浮点数
	fmt.Printf("f1:%F\n", f1)    // f1:6.123456 浮点数 %f 的同义词。
	fmt.Printf("f1:%0.2f\n", f1) // f1:6.12 浮点数 保留两位小数
	fmt.Printf("f1:%e\n", f1)    // f1:6.123456e+00 浮点数 科学计数法 小写e
	fmt.Printf("f1:%E\n", f1)    // f1:6.123456E+00 浮点数 科学计数法 大写E
	fmt.Printf("f1:%g\n", f1)    // f1:6.123456 浮点数 对于较大的指数使用 %e，否则使用 %f
	fmt.Printf("f1:%G\n", f1)    // f1:6.123456 浮点数 对于较大的指数使用 %E，否则使用 %F
	fmt.Printf("f1:%x\n", f1)    // f1:6.123456 浮点数 以十六进制表示，带有十进制的指数部分
	fmt.Printf("f1:%X\n", f1)    // f1:6.123456 浮点数 大写字母的十六进制表示，例如 -0X1.23ABCP+20

	fmt.Println("=======================")

	// 接下来是布尔类型的格式化打印的占位符示例
	b1 := true
	fmt.Printf("b1:%t\n", b1) // b1:true

	fmt.Println("=======================")

	// 接下来是字符串和字节数组的格式化打印的占位符示例
	name := "MagicGopher"
	byteArray := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67}
	fmt.Printf("name:%s\n", name)                 // name:MagicGopher
	fmt.Printf("name:%q\n", name)                 // name:"MagicGopher"
	fmt.Printf("byteArray字节数组内容:%x\n", byteArray) // 以十六进制（base 16）形式打印字节切片或字符串，其中每个字节用两个小写字母表示。
	fmt.Printf("byteArray字节数组内容:%X\n", byteArray) // 以十六进制（base 16）形式打印字节切片或字符串，其中每个字节用两个大写字母表示。

	fmt.Println("=======================")

	// 接下来是切片类型和指针类型的格式化打印的占位符示例
	var slice = []int{1, 2, 3, 4, 5}
	var pointer = &slice
	fmt.Printf("slice:%v\n", slice)     // slice:[1 2 3 4 5]
	fmt.Printf("pointer:%p\n", pointer) // pointer:[计算机为这个指针地址分配的内存地址]
}
