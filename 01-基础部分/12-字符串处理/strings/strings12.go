package main

import "strings"

func main() {
	// TrimPrefix函数的作用：删除字符串s的前缀字符串prefix
	// 去除字符串开头的指定前缀
	str1 := "https://example.com/path/to/resource"
	prefix := "https://"
	trimmed1 := strings.TrimPrefix(str1, prefix)
	println(trimmed1)
	// 当前缀不存在时,返回原字符串
	str2 := "example.com/path/to/resource"
	trimmed2 := strings.TrimPrefix(str2, prefix)
	println(trimmed2)
}
