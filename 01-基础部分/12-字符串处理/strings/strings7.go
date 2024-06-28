package main

import (
	"fmt"
	"strings"
)

func main() {
	// 包含某个字符做为后缀
	str := "hello world"
	fmt.Println(strings.HasSuffix(str, "ld"))  // true
	fmt.Println(strings.HasSuffix(str, "wor")) // false
}
