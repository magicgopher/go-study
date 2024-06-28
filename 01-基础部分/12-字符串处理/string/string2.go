package main

import "fmt"

func main() {
	str1 := "hello\aworld" // \a 响铃
	fmt.Println(str1)
	str1 = "hello\bworld" // \b 退格
	fmt.Println(str1)
	str1 = "hello\fworld" // \f 换页
	fmt.Println(str1)
	str1 = "hello\nworld" // \n 换行
	fmt.Println(str1)
	str1 = "hello\rworld" // \r 回车
	fmt.Println(str1)
	str1 = "hello\tworld" // \t 制表符
	fmt.Println(str1)
	str1 = "hello\vworld" // \v 纵向制表符
	fmt.Println(str1)
	str1 = "hello\"world" // \ 斜杠
	fmt.Println(str1)
	str1 = "打印字母 'A' 的十六进制转义字符:\x41"
	fmt.Println(str1)
	str1 = "打印希腊字母 'α' 的十六进制转义字符:\u03b1"
	fmt.Println(str1)
	str1 = "打印一个自定义的十六进制字符:\xff"
	fmt.Println(str1)
	str1 = "打印字母 'A' 的八进制转义字符:\101"
	fmt.Println(str1)
	str1 = "打印希腊字母 'α' 的八进制转义字符:\316\261"
	fmt.Println(str1)
	str1 = "打印一个自定义的八进制字符:\377"
	fmt.Println(str1)
}
