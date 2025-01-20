package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os 包的 Open 函数打开文件
	file, err := os.Open("README.md")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	// 打印文件名称和打印文件打开成功
	fmt.Println(file.Name(), "open file success.")
}
