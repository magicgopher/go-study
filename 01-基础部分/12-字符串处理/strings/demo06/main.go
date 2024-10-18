package main

import (
	"fmt"
	"strings"
)

func main() {
	// 包含某个字符做为前缀
	str := "hello world!"
	fmt.Println(strings.HasPrefix(str, "he"))  // true
	fmt.Println(strings.HasPrefix(str, "llo")) // false
}
