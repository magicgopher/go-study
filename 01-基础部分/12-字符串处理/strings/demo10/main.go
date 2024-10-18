package main

import (
	"fmt"
	"strings"
)

func main() {
	// Trim()函数的作用：去除字符串首尾的指定字符
	str1 := "   hello, world!   "
	trimmed1 := strings.Trim(str1, " ")
	fmt.Println(trimmed1) // 输出: hello, world!
	str2 := "---hello, world!---"
	trimmed2 := strings.Trim(str2, "-")
	fmt.Println(trimmed2) // 输出: hello, world!
	str3 := "***hello, world!***"
	trimmed3 := strings.Trim(str3, "*")
	fmt.Println(trimmed3)              // 输出: hello, world!
	trimmed4 := strings.Trim(str1, "") // cutset 为空字符串时,Trim() 函数不会做任何修改
	fmt.Println(trimmed4)
}
