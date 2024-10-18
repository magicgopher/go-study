package main

import "fmt"

func main() {
	// Scanln 函数示例
	var name string
	var age int

	fmt.Println("请输入姓名和年龄(以空格分隔):")
	n, err := fmt.Scanln(&name, &age)
	if err != nil {
		fmt.Println("输入格式有误:", err)
		return
	}
	fmt.Printf("成功读取了 %d 个值:\n", n)
	fmt.Printf("姓名: %s, 年龄: %d\n", name, age)
}
