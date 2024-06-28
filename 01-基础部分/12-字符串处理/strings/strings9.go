package main

import (
	"fmt"
	"strings"
)

func main() {
	// ToLower函数：将字符串转换为小写
	str1 := strings.ToLower("GOLANG")
	fmt.Printf("str1:%s\n", str1) // str1:golang
	str2 := strings.ToLower("你好") // 不支持中文
	fmt.Printf("str2:%s\n", str2) // str2:你好
	// ToUpper函数：将字符串转换为大写
	str3 := strings.ToUpper("java")
	fmt.Printf("str3:%s\n", str3) // str3:JAVA
}
