package main

import (
	"fmt"
	"strings"
)

func main() {
	// Replace()函数 用来替换字符串中的内容
	// 参数1: 字符串 参数2: 要替换的内容 参数3: 替换的内容 参数4: 替换的次数
	str1 := "Hello, World!"
	newStr := strings.Replace(str1, "World", "Go", 1)
	fmt.Println(newStr)
	// ReplaceAll()函数 用来替换字符串中的内容
	// 参数1: 字符串 参数2: 要替换的内容 参数3: 替换的内容
	newStrAll := strings.ReplaceAll(str1, "o", "O")
	fmt.Println(newStrAll)
}
