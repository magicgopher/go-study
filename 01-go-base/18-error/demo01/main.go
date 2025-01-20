package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os 包的 Open 函数打开文件
	file, err := os.Open("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 打印文件名称和打印文件打开成功
	fmt.Println(file.Name(), "open file success.")
}
