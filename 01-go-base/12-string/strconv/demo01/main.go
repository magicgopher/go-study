package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串与整数之间的转换
	// Atio函数：将字符串转换为整数
	r1, _ := strconv.Atoi("12345")
	fmt.Printf("r1类型:%T,r1值:%v\n", r1, r1)
	// Iota函数：将整数转换为字符串
	r2 := strconv.Itoa(54321)
	fmt.Printf("r2类型:%T,r2值:%v\n", r2, r2)
}
