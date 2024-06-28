package main

import (
	"fmt"
	"strings"
)

func main() {
	// 定义字符串
	str1 := "a,b,c"
	// Split()函数作用：将字符串 s 按照指定的分隔符 sep 切分成一个字符串切片。
	parts := strings.Split(str1, ",") // [a b c]
	fmt.Println(parts)
	// SplitN()函数作用：将字符串 s 按照指定的分隔符 sep 切分成最多 n 个子串的字符串切片,不会忽略分隔符。
	partsN := strings.SplitN(str1, ",", 2)
	fmt.Println(partsN)
}
