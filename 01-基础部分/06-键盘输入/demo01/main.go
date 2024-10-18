package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 使用bufio包创建Reader实例读取键盘输入内容
	// os.Stdin全称"standard input"。
	reader := bufio.NewReader(os.Stdin)
	// 读取
	fmt.Println("请输入姓名")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 打印输出
	fmt.Println("Name:", name)

}
