package main

import (
	"fmt"
	"strings"
)

func main() {
	// TrimSuffix函数的作用：删除字符串末尾的指定后缀
	// 去除字符串末尾的指定后缀
	str1 := "file.txt"
	suffix := ".txt"
	trimmed1 := strings.TrimSuffix(str1, suffix)
	fmt.Println(trimmed1)
	// 当后缀不存在时,返回原字符串
	str2 := "file.zip"
	trimmed2 := strings.TrimSuffix(str2, suffix)
	fmt.Println(trimmed2)
}
