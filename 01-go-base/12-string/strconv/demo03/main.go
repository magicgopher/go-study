package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串与浮点数之间的转换
	// ParseFloat函数的作用：将字符串转换为浮点数
	f1, _ := strconv.ParseFloat("3.1415", 64) // 3.1415, nil
	fmt.Printf("f1类型:%T,f1值:%v\n", f1, f1)
	f2 := strconv.FormatFloat(9.8765, 'f', -1, 64) // "9.8765"
	fmt.Printf("f2类型:%T,f2值:%v\n", f2, f2)
}
