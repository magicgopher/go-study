package main

import "fmt"

func main() {
	// 使用初始化语句
	// error错误处理，后面文档中有介绍
	if x, err := getData(); err == nil {
		fmt.Println("数据获取成功:", x)
	} else {
		fmt.Println("数据获取失败:", err)
	}
}

// getData 获取数据
func getData() (int, error) {
	return 42, nil
}
