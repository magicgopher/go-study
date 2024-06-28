package main

import (
	"fmt"
	"strings"
)

func main() {
	// Repeat()函数作用：创建一个新的字符串，由指定的字符串 s 重复 count 次拼接而成。
	// 复制几次字符串
	str := "hello"
	fmt.Println(strings.Repeat(str, 2)) // hellohello
}
