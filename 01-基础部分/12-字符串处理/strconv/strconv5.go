package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串与无符号整数之间的转换
	// ParseUint函数的作用：将字符串转换为无符号整数
	u1, _ := strconv.ParseUint("123", 10, 64) // 10表示进制，64表示转换后的类型
	fmt.Printf("u1类型:%T,u1值:%v\n", u1, u1)
	u2 := strconv.FormatUint(4321, 10) // "4321"
	fmt.Printf("u2类型:%T,u2值:%v\n", u2, u2)
}
