package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串查找
	str := "Hello, World!"
	index := strings.Index(str, "World")     // 7
	lastIndex := strings.LastIndex(str, "o") // 8
	fmt.Printf("index:%d, lastIndex:%d\n", index, lastIndex)
}
