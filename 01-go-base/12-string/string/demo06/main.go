package main

import "strings"

func main() {
	// 使用 strings.Builder 构建字符串
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString("World")
	builder.WriteString("Go语言学习笔记")
	// builder.String()方法返回构造好的字符串
	result := builder.String()
	println(result)
}
