package main

import (
	"fmt"
)

func main() {
	// 字符串拼接，只能是字符串类型(其他数据类型需要转换成字符串类型再进行拼接)
	str1 := "hello"
	str2 := "world"
	r1 := str1 + str2
	fmt.Printf("r1类型:%T,r1的值:%s\n", r1, r1)
}
