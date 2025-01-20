package main

import (
	"fmt"
)

func main() {
	// 字符串比较
	s1 := "hello"
	s2 := "hello"
	s3 := "world"
	fmt.Println(s1 == s2) // true
	fmt.Println(s1 == s3) // false
	fmt.Println(s1 != s2) // false
	fmt.Println(s1 != s3) // true
}
