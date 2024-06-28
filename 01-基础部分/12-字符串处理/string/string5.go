package main

import (
	"fmt"
	"strings"
)

func main() {
	// 使用strings包的Join()函数来拼接字符串
	parts := []string{"hello", "world", "Java", "Go", "Python"}
	// elems []string: 字符切片
	// sep string: 这是一个字符串,表示用于分隔 elems 中各个字符串元素的分隔符.
	result := strings.Join(parts, "+")
	fmt.Println(result)
}
