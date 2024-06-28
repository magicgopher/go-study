package main

import (
	"fmt"
	"strings"
)

func main() {
	// Count()函数作用：统计字符串中某个字符出现的次数
	// 统计字符串中每个字符出现的次数
	str := "hello world"
	for _, v := range str {
		// 统计字符出现的次数
		count := strings.Count(str, string(v))
		fmt.Printf("%c: %d\n", v, count)
	}
}
