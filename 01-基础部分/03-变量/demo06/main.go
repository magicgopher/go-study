package main

import "fmt"

// Go语言变量注意事项5
// 这里尝试使用简短变量声明定义全局变量，会导致编译错误
// age := 26 // 简短定义全局变量
var age = 26

func main() {
	fmt.Println("Age:", age)
}
