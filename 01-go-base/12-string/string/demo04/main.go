package main

import "fmt"

func main() {
	// 使用fmt.Sprintf()函数来进行字符串拼接
	var str = "hello"
	var str2 = "world"
	// 返回的是按照格式拼接好的字符串
	info1 := fmt.Sprintf("%s %s", str, str2)
	fmt.Println(info1)
	info2 := fmt.Sprintf("%s , %s", str, str2)
	fmt.Println(info2)
}
