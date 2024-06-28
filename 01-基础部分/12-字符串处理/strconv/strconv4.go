package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串与整数之间的转换
	i1, _ := strconv.ParseInt("123", 10, 64) // 123, nil
	fmt.Printf("i1类型:%T,i1值:%v\n", i1, i1)
	i2 := strconv.FormatInt(321, 2) // 二进制 321
	fmt.Printf("i2类型:%T,i2值:%v\n", i2, i2)
	i3 := strconv.FormatInt(321, 8) // 八进制 321
	fmt.Printf("i3类型:%T,i3值:%v\n", i3, i3)
	i4 := strconv.FormatInt(321, 10) // 十进制 321
	fmt.Printf("i4类型:%T,i4值:%v\n", i4, i4)
	i5 := strconv.FormatInt(321, 16) // 十六进制 321
	fmt.Printf("i5类型:%T,i5值:%v\n", i5, i5)
}
