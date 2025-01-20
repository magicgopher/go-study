package main

import "fmt"

func main() {
	// 类型 switch语句示例
	// switch语句 case后面还可以跟表达式，switch后面就不能有表达式。
	score := 85
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
		break // 跳出当前 switch 语句
	case score >= 70:
		fmt.Println("中等")
	default:
		fmt.Println("不及格")
	}
}
