package main

import "fmt"

func main() {
	// len函数获取字符串长度
	s1 := "hello world"
	// 这个字符串包含 12 个字符(包括空格)
	// 在 UTF-8 编码下,每个字符占用 1 个字节
	// 因此这个字符串的总大小为 12 个字节
	fmt.Println(len(s1))
}
