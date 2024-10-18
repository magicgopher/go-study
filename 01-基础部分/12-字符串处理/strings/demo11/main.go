package main

import (
	"fmt"
	"strings"
)

func main() {
	// TrimSpace()函数的作用：去除字符串首尾的空格，并返回处理后的字符串。
	str1 := "   hello, world!   "
	trimmed1 := strings.TrimSpace(str1) // 输出："hello, world!"
	fmt.Println(trimmed1)
	str2 := "\t\nhello, world!\n\t"
	trimmed2 := strings.TrimSpace(str2) // 输出："hello, world!"
	fmt.Println(trimmed2)
	str3 := "hello, world!"
	trimmed3 := strings.TrimSpace(str3) // 输出："hello, world!"
	fmt.Println(trimmed3)
}
