package main

import "fmt"

func main() {
	// Scanf函数示例

	var name string
	var age int

	// Scanf函数读取用户输入
	fmt.Println("请输入姓名和年龄,格式: name:李四 age:20")
	n, err := fmt.Scanf("name:%s age:%d", &name, &age) // 键盘输入格式: name:张三 age:18
	if err != nil {
		fmt.Println("输入格式有误:", err)
		return
	}
	fmt.Printf("成功读取了 %d 个值:\n", n)
	// 格式化打印
	fmt.Printf("姓名:%s, 年龄:%d\n", name, age)
}
